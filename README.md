# go-movefilesoutside
just a simple script to move all file inside folders to outside folders

## Usage

Put your folder with many files to "FILES" folder then run following command:

```bash
go run main.go -mode MODE -in FOLDERNAME -out FOLDERNAME -limit NUMBER
```
Flag
<ul>
<li>-mode : select between move or copy (default)</li>
<li>-in : folder to scan</li>
<li>-out : folder to store result</li>
<li>-limit : max files scanned to execute</li>
</ul>