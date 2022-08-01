# EGoVault
![chest](EGoVault.png?raw=true "EGoVault")

[About_EGoVault]
- EGoVault allows write, append info into a specific file, using enclave ciphering operations.

[Configure]
- Set your source and target paths into `enclave.json`
- Configure user db path into userOperations.go `DB_PATH` variable

[Compile] 
- make clean && make

[ARGUMENTS]
- --help Show helps menu
- --seal Specify message and path where you can save the encrypted message
- --unseal Specify path of the seal file
- --append Append new data to a existent file

[EXAMPLES]
- ego run app --seal "your message using double quotes" $(pwd)/db/your_secret_file
- ego run app --unseal $(pwd)/db/your_secret_file
- ego run app --append "your NEW message using double quotes" $(pwd)/db/your_secret_file