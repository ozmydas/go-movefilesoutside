# go-movefilesoutside
just a simple script to move all file inside folders to outside folders

## Usage

Put your folder with many files to "FILES" folder then run following command:

```bash
go run main.go -opt MODE -in FOLDERNAME -out FOLDERNAME
```
Flag
-opt : select between move or copy (default)
-in : folder to scan
-out : folder to store result