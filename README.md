# EGoVault

[ARGUMENTS]
--help\t\tShow helps menu
--seal\t\tSpecify message and path where you can save the encrypted message
--unseal\tSpecify path of the seal file
--append\tAppend new data to a existent file

[EXAMPLES]
ego run app " + --seal \"your message using double quotes\" <filePathDestination>"
ego run app " + --unseal <filePathDestination>
ego run app " + --append \"your NEW message using double quotes\" <existentFile>