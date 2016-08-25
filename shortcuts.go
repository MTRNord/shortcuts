package main

import (
   "os"
   "fmt"
   "flag"
   "unsafe"
   "path/filepath"
)

var (
   version string = "v2.0.1-beta"
   vers bool
   file    string
)

func init() {
   flag.StringVar(&file, "file", "false", "File to find the distance between.")
   flag.BoolVar(&vers, "version", false, "[Optional] Return version of bindist.")   
}

//return: found, off1, off2, errors
func handleFile(fp *os.File) {


}

//callback for walk needs to match the following:
//type WalkFunc func(path string, info os.FileInfo, err error) error
func readFile (path string, fi os.FileInfo, err error) error {
   
   f, err := os.Open(path)
   defer f.Close()   //closing the file
   if err != nil {
      fmt.Fprintln(os.Stderr, "ERROR:", err)
      os.Exit(1)  //should only exit if root is null, consider no-exit
   }

   switch mode := fi.Mode(); {
   case mode.IsRegular():
      handleFile(f)
   case mode.IsDir():
      fmt.Fprintln(os.Stderr, "INFO:", fi.Name(), "is a directory.")      
   default: 
      fmt.Fprintln(os.Stderr, "INFO: Something completely different.")
   }
   return nil
}

func processfiles() {
   filepath.Walk(file, readFile)
}

func main() {

   flag.Parse()
   var verstring = "shortcutz version"
   if vers {
      fmt.Fprintf(os.Stderr, "%s %s \n", verstring, version)
      os.Exit(0)
   } else if flag.NFlag() <= 2 {    // can access args w/ len(os.Args[1:]) too
      fmt.Fprintln(os.Stderr, "Usage:  shortcutz [-file ...]")
      fmt.Fprintln(os.Stderr, "                  [Optional -version]")
      fmt.Fprintln(os.Stderr, "")
      fmt.Fprintln(os.Stderr, "Output: [STRING] TBD Structure of some sort...")
      fmt.Fprintf(os.Stderr, "Output: [STRING] '%s ...'\n\n", verstring)
      flag.Usage()
      os.Exit(0)
   }

   var x = 1
   if x == 2 {
      fmt.Println("4C0000000114020000000000C000000000000046")

      var ross ShellLinkHeader
      const infoSize = unsafe.Sizeof(ross)

      fmt.Println(infoSize)      
   }

}
   