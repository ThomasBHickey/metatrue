metatrue
========

Merge of TrueType and MetaFont

The impetus for this code is to revive some old MetaFont79 code.  The goal is
a system with the power of MetaFont that understands and generates good TrueType
fonts.

As a first exercise I'm attempting to rewrite the 2014 MetaFont in Go.

This is not the first time I have ported a MetaFont, as I once rewrote
the original MetaFont SAIL code into a working version on
Tandem computers in the TAL language (and used it to start our font work
at OCLC).  See Hickey, Thomas B. The status of Metafont at OCLC pp. 35–38 at
https://www.tug.org/TUGboat/tb02-2/tb03fonts.pdf 
  
The plan as of September 2014 is to follow Knuth's structure, but along the way use features
of Go (golang) as they make sense.  In particular I am not planning on using
the existing memory management and string routines in MetaFont.
