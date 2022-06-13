# Avatar-me
Given personal information, such as an email address, an IP address or a public key, a unique avatar is generated.
(Repository belonging to the Golang course)

## How to
The information must be entered as a string to TheInfo variable:

avatar.TheInfo = "This is an example"

Next is running DefaultFeaturesGeneration, function wich provides a pointer to a struct containing the interfaces
needed, and an error if it occurs:

utilitiesNeeded , err := avatar.DefaultFeaturesGeneration()

Finally the identicon can be made and saved as an .png image, named after string contained in avatar.TheInfo using
GenerateAndSaveAvatar() that return an error, if it occurs:

err = aver.GenerateAndSaveAvatar()