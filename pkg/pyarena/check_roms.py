#!/usr/bin/env python
import diambraArena, os, sys

for arg in sys.argv[1:]:
    diambraArena.checkGameSha256(os.path.join(os.getenv('DIAMBRAROMSPATH'), arg))