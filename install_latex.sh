#!/bin/bash
tlmgr install ifetex abntex2 enumitem lipsum breakurl collection-fontsrecommended
tlmgr update l3kernel

# then just run xelatex file.tex