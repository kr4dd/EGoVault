# EGoVault

[ARGUMENTS]
- --help Show helps menu
- --seal Specify message and path where you can save the encrypted message
- --unseal Specify path of the seal file
- --append Append new data to a existent file

[EXAMPLES]
- ego run app --seal "your message using double quotes" "filePathDestination"
- ego run app --unseal "filePathDestination"
- ego run app --append "your NEW message using double quotes" "existentFile"