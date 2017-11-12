# grep utility #
Search for a string in a given filepath 

# todo #
- Recursive search might need recursion
- refactor main part to function to support recursion

# usage #
- Input file path and string to search
- Output file names which contain the string and match line

# Notes #
- Have to check for file type to detect binary files using http.DetectContentType function
- strings.Index is doing all the work