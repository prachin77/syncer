package files

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	"github.com/prachin77/pkr/models"
)

var (
	sending_workspace     = models.SendWorkSpaceFolder{}
	recieved_workspace    = models.RecievedWorkSpaceFolder{}
	workspace_path        string
	workspace_hosted_date string
	workspace_hosted_port string
	username              string
)

func GetIpAdd() (string, error) {
	cmd := exec.Command("ipconfig")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Convert output to string
	outputStr := string(output)

	// Extract IPv4 address (basic regex for Windows `ipconfig` output)
	re := regexp.MustCompile(`IPv4 Address[.\s]*:\s*([\d.]+)`)
	match := re.FindStringSubmatch(outputStr)
	if len(match) < 2 {
		return "", fmt.Errorf("could not find IP address in ipconfig output")
	}

	return match[1], nil
}

func GetClientPublicKeyFilepath() string {
	my_public_key_filepath, err := filepath.Abs("./config/publickey.pem")
	if err != nil {
		fmt.Println("error retrieving client public key file path !")
		return ""
	} else {
		return my_public_key_filepath
	}
}

func GetHostWorkspaceInfo(decrypted_workspace_password string, workspace_name string) (string, string, string, string, string) {
	file, err := os.Open(models.USER_CONFIG_FILE)
	if err != nil {
		fmt.Println("failed to open host user config file : ", err)
		return "", "", "", "", ""
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&user_config)
	if err != nil {
		fmt.Println("error decoding json data from host user config file : ", err)
		return "", "", "", "", ""
	}

	for _, workspace := range user_config.SendWorkSpaces {
		if workspace.Workspace_Name == workspace_name && workspace.Workspace_Password == decrypted_workspace_password {
			username = user_config.Username
			workspace_hosted_port = user_config.Port
			workspace_path = workspace.Workspace_Path
			workspace_hosted_date = workspace.Workspace_Hosted_Date
			return workspace_path, workspace_name, workspace_hosted_date, workspace_hosted_port, username
		}
	}

	fmt.Println("no matching workspace found or invalid credentials")
	return "", "", "", "", ""
}

// GPT code but understand later
func ZipData(workspacePath string, workspaceName string) (string, error) {
	// Create the name of the zip file
	ZipFolderName := workspaceName + "_zip.zip"

	// Attempt to create the zip file on the disk. If this fails, return an error.
	zipFile, err := os.Create(ZipFolderName)
	if err != nil {
		return "", fmt.Errorf("error creating zip file: %v", err)
	}
	defer zipFile.Close()

	// Create a new zip writer to write compressed data into the zip file.
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	var fileCount, folderCount int
	var fileNames, folderNames []string

	// Ensure the workspace path ends with the correct slash for consistent path calculations.
	workspacePath = filepath.Clean(workspacePath) + string(os.PathSeparator)

	// Walk through the workspace directory to process each file and folder.
	err = filepath.Walk(workspacePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip adding the zip file itself to avoid recursive inclusion.
		if filepath.Base(path) == ZipFolderName {
			return nil
		}

		// Skip the root workspace directory itself.
		if path == workspacePath {
			return nil
		}

		// Exclude the "config" directory and all its contents from the zip.
		if info.IsDir() && info.Name() == "config" {
			return filepath.SkipDir
		}

		// Exclude files with the ".exe" extension.
		if !info.IsDir() && filepath.Ext(info.Name()) == ".exe" {
			return nil
		}

		// Calculate the relative path of the current file or folder from the workspace root.
		relPath, err := filepath.Rel(workspacePath, path)
		if err != nil {
			return fmt.Errorf("error calculating relative path: %v", err)
		}

		// Create a header for the current file or folder to add it to the zip.
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return fmt.Errorf("error creating file header for %s: %v", path, err)
		}

		// Set the header name to the relative path, ensuring compatibility across platforms.
		header.Name = filepath.ToSlash(relPath)

		// If the current item is a directory, ensure the header name ends with a slash.
		if info.IsDir() {
			header.Name += "/"
			folderCount++
			folderNames = append(folderNames, header.Name)
		} else {
			// For files, set the compression method to Deflate.
			header.Method = zip.Deflate
			fileCount++
			fileNames = append(fileNames, header.Name)
		}

		// Add the header to the zip writer.
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("error adding header to zip for %s: %v", path, err)
		}

		// Skip writing content for directories, as they don't have data.
		if info.IsDir() {
			return nil
		}

		// Open the current file for reading its content.
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("error opening file %s: %v", path, err)
		}
		defer file.Close()

		// Copy the file's content into the zip writer.
		_, err = io.Copy(writer, file)
		if err != nil {
			return fmt.Errorf("error writing file %s to zip: %v", path, err)
		}

		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error walking through workspace directory: %v", err)
	}

	// Print summary of the zipping process
	fmt.Printf("Workspace successfully zipped as %s\n", ZipFolderName)
	fmt.Printf("Folders zipped: %d\n", folderCount)
	fmt.Printf("Files zipped: %d\n", fileCount)
	fmt.Printf("Folders zipped names: %v\n", folderNames)
	fmt.Printf("Files zipped names: %v\n", fileNames)

	// Return the path of the created zip file
	return filepath.Abs(ZipFolderName)
}

func UnZipData(ZipFilePath string, DestPath string) error {
	zipReader, err := zip.OpenReader(ZipFilePath)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer zipReader.Close()

	var fileCount, folderCount int

	for _, file := range zipReader.File {
		fullPath := filepath.Join(DestPath, file.Name)

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
			folderCount++
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create parent directory: %v", err)
		}

		fileReader, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file in zip: %v", err)
		}
		defer fileReader.Close()

		outFile, err := os.Create(fullPath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer outFile.Close()

		if _, err := io.Copy(outFile, fileReader); err != nil {
			return fmt.Errorf("failed to copy file contents: %v", err)
		}

		fileCount++
	}

	fmt.Printf("%d] files unzipped  \n%d] folders unzipped\n", fileCount, folderCount)
	return nil
}

func SaveDataToZip(data []byte, zipFilePath string) error {
	zippedFile, err := os.Create(zipFilePath)
	if err != nil {
		fmt.Println("error creating data from host zip file to client zip file !")
		return err
	}
	defer zippedFile.Close()

	zippedFile.Write(data)

	return nil
}

func WriteRecivedWorkspaceInConfigFile(workspace_name string, workspace_path string, workspace_IP string) error {
	currentDate := time.Now().Format("2006-01-02 15:04:05")

	data, err := os.ReadFile(models.USER_CONFIG_FILE)
	if err != nil {
		fmt.Println("error reading user config file : ", err)
		return err
	}

	if err := json.Unmarshal(data, &user_config); err != nil {
		fmt.Println("error unmarshalling data in JSON format : ", err)
		return err
	}

	recieved_workspace.Recieved_Date = currentDate
	recieved_workspace.Workspace_Name = workspace_name
	recieved_workspace.Workspace_IP = workspace_IP
	recieved_workspace.Workspace_Path = workspace_path

	user_config.RecievedWorkspaces = append(user_config.RecievedWorkspaces, recieved_workspace)

	updatedData, err := json.MarshalIndent(user_config, "", " ")
	if err != nil {
		fmt.Println("error marshalling updated data into JSON format : ", err)
		return err
	}

	err = os.WriteFile(models.USER_CONFIG_FILE, updatedData, os.ModePerm)
	if err != nil {
		fmt.Println("error writing data into user config file : ", err)
	}

	fmt.Println("Data succesfully written into file ...")

	return nil
}
