CRTD to MSL
===

Converts CANBUS log from CRTD format to something megalogviewer can open

DBC file from http://www.msextra.com/doc/pdf/Megasquirt_CAN-2014-10-27.dbc
Field "map" renamed to "map1" to avoid clobbering the golang map builtin
Does not generate on windows with https://github.com/einride/can-go DBC code generator due to backslash file paths 


