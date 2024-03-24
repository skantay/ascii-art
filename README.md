# ASCII-art
```
@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
""WW^^@@@@^^``````^^BB@@^^``^^&&``
^^&&^^@@````^^``&&``@@````^^^^&&``
``WW&&^^""``^^WW&&&&""``^^^^&&88``
^^8888&&&&&&WW88&&88WW&&&&88&&WW``
@@``&&88888888WW&&WW88&&88WW88^^$$
@@""88&&&&&&&&888888&&``^^&&88``$$
@@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
@@@@@@^^888888&&88&&&&MM88^^BB$$$$
@@@@@@BB````&&&&&&&&88""``BB@@BB$$
$$@@$$$$$$$$``````````@@$$@@$$$$$$
```

The `asciiart` package in Go provides a straightforward and efficient solution for converting regular text into captivating ASCII art representations. [ASCII](https://en.wikipedia.org/wiki/ASCII_art) art is a technique that uses printable characters from the ASCII standard to create visual images or designs. With this package, developers can effortlessly integrate ASCII art creation into their applications, adding a nostalgic or creative touch to text-based interfaces, command-line tools, or even web applications.

## Usage

To convert import package
```
github.com/skantay/asciiart
```

And provide a text and file to function `GetASCII`

```go
package main

import "github.com/skantay/asciiart"

func main() {
    file, _ := os.Open("font.txt")

    asciiArt, _ := asciiart.GetASCII("hello", file)
}
```

File must contain font, and chars from 32nd char till 127th of [ASCII table](https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html). 

- Requirements for font

1. Height of char must be 8
2. Chars are seperated with a new line
3. Skip first line of file
4. IMPORTANT!!! Chars must be ordered from 32nd char till 127th of [ASCII table](https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html)

See an example in banners folder, and feel free to download them.

## Examples

```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```