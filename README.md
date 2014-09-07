metatrue
========

Merge of TrueType and MetaFont

The impetous for this code is to revive some old MetaFont79 code.  The goal is
a system with the power of MetaFont that understands and generates good TrueType
fonts.

As a first excercise I'm attempting to rewrite the 2014 MetaFont in Go.

This is not the first time I have porated a MetaFont, as I once rewrote
the original MetaFont SAIL code into a working verion on
Tandem computers in the TAL language (and used it to start our font work
at OCLC).

The plan is to follow Knuth's structure, but along the way use features
of Golang as they make sense.  In particular I am not planning on using
the existing memory management and string routines in MetaFont.
