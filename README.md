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


Update Sept 28th:  I'm probably going to put this on hold for a while, although
reading some about TeX makes it clear that the memory limitations of the current
system are causing problems.  I think there is an approach that might not be too
hard to implement.  I'd separate the two sections of memory that grow together.
Each of them would be a slice of pointers through a Node interface to various
node types.  That way we could still use integers (of some size) as pointers
to the nodes (actually array/slice indices to obtain pointers to the nodes).
I believe that would greatly simplify the port.  I'd also like to make it possible
to avoid limitations like not being able to enter 4096 as a number.  Looks like
that is (probably deliberately) coded in lots of places.  Probably not that
hard to eliminate, but I expect that makes it not a compliant MetaFont at that
point.

October 2014: Still working on it!  Currently working trying to get the parser to actually parse something!
