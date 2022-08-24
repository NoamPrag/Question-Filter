# Question-Filter
A tool to filter the results of survey responses from patients.

The program receives an Excel file (.xlsx) with the survey results, and spits out a new Excel file, including only the questions specified by the user.

## Usage
---
### How to run:
1. Run the program
2. Select your input Excel file
3. Enter "filters" for events, forms and question indices
4. Save the filtered Excel file

### The "filters":
A filter could be just a number, a range of numbers, or a combination of both, separated by commas.
A range is denoted by two values and a hyphen (-) in between.
#### Example:
Filtering questions using "1, 3-5, 8" will select all the first question, all the questions that have an index between 3 and 5, and the 8th question. <br/>
* Note that white space doesn't matter, so "1,3-5,8" will be the same as the example above.


## How to download
---
1. Clone the repository (press the green "code" button, download zip and then extract the files)
2. Inside the "app" directory you'll see "macos" and "windows" folders containing the application for the two operating systems.

### Build from source
If you wish to alter the functionality of the program, change to code and then run the "build.sh" script. This script will compile the application for mac and windows and will place it in the app directory instead of the current build.
* Note that this script is currently suitable for running on macos only.

## Tools used
---
The program is written in the Go programming language, and makes use of a couple of libraries:
1. Zenity - for user interface and input
2. Excelize - for interacting with Excel files