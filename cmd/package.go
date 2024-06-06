package cmd

type Package struct {
	MFile MFile
	// Readme
	// Icon
	// Deps?
	// outputFile??
	// If readme exists, will replace MFile description

	// compile to ./build
	// config.lua
}

/*
Package Structure


/
mfile
readme.md
.gitignore
.gitattributes
Taskfile.yml
src/
	resources/   <--- This is moved raw. No PKGNAME replace, etc.
*/
