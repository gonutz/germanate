Trying to write German on an English keyboard proves to be hard.
Entering umlauts like ä and ß is not easy.

This program will insert umlauts in your clipboard.
You can copy your text, run this program, then paste your German text.

Replacements done by this program:

	Ae  - Ä
	ae  - ä
	Oe  - Ö
	oe  - ö
	Ue  - Ü
	ue  - ü
	sss - ß

Note that you need 3 s for ß. While it is natural to write ae, oe and ue,
it is usually ss that is being replaced with ß. But ss is also valid in
German so if we replace all ss with ß that would be wrong. Hence this
quirk.

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

