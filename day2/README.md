# grep utility #
Search for a string in a given path

# todo #
- Recursive search might need recursion
- refactor main part to function to support recursion

# Notes #
- Have to check for file type to detect binary files using http.DetectContentType function
- strings.Index is doing all the work