 HTML Link Parser

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/link)

## Exercise details

This is a package that makes it easy to parse an HTML file and extract all of the links (`<a href="">...</a>` tags). For each extracted link it returns a data structure that includes both the `href`, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

