package content

import "fmt"

var opfTmpl = `<?xml version='1.0' encoding='utf-8'?>
<package xmlns="http://www.idpf.org/2007/opf" version="2.0" unique-identifier="%s">
  <metadata>
    <dc-metadata xmlns:dc="http://purl.org/dc/elements/1.1/">
      <dc:title>%s</dc:title>
      <dc:language>en-gb</dc:language>
      <meta content="cover-image" name="cover"/>
      <dc:creator>%s</dc:creator>
      <dc:publisher>%s</dc:publisher>
      <dc:subject>News</dc:subject>
      <dc:date>2011-05-09</dc:date>
      <dc:description>%s</dc:description>
    </dc-metadata>
    <x-metadata>
      <output content-type="application/x-mobipocket-subscription-magazine" encoding="utf-8"/>
    </x-metadata>
  </metadata>
  <manifest>
  %s
  </manifest>
  <spine toc="nav-contents">
  %s
  </spine>
  <guide>
    <reference href="contents.html" type="toc" title="Table of Contents"/>
  </guide>
</package>
`

//GenerateOpf ...
func GenerateOpf(feed Feed, manifest string, spine string) string {
	id := feed.ID
	title := feed.Title
	creator := ""
	publisher := ""
	description := ""

	return fmt.Sprintf(opfTmpl,
		id,
		title,
		creator,
		publisher,
		description,
		manifest,
		spine,
	)
}