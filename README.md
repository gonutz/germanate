Trying to write German on an English keyboard proves to be hard.
Entering umlauts like ä and ß is not easy.

This program will insert umlauts in your clipboard.
You can copy your text, run this program, then paste your German text.

Replacements done by this program:

	AE - Ä
	aE - ä
	OE - Ö
	oE - ö
	UE - Ü
	uE - ü
	sS - ß

Note that the replacement only happens when the second letter is upper-case.
Except in abbreviations, this should rarely happen in regular text. We cannot
use the more common way of writing for example `ue` for `ü` because words like
`neue` would turn into `neü` which is not what we want.

# Installation

	go install github.com/gonutz/germanate@latest

# Intrgration with Ubuntu

For easy access I added this program to my favorites in Ubuntu.
This way it appears in my search results and I have an icon for it on
the launch bar. I can thus either:

 - Press the super key and search for it by typing germanate
 - Click the favorite icon with the mouse
 - Hit super key + 5 (it is my 5th icon)

to translate my texts.

To add it you must create this file

	/usr/share/applications/germanate.desktop

with this content

	[Desktop Entry]
	Type=Application
	Encoding=UTF-8
	Name=Germanate
	Comment=Translate German umlauts in the clipboard
	Exec=/home/me/go/bin/germanate
	Icon=/home/me/go/bin/germanate.png
	Terminal=false

Modify the Exec path accordingly. The Icon is optional and not provided
by this repo. You can remove that line or create an icon yourself.

Now make the file executable

	sudo chmod +x /usr/share/applications/germanate.desktop

After this you can search for Germanate, right click the search result
icon and say Add to Favorites.

