package po

import "embed"

//go:generate sh -c "find ../src -name '*.go' -o -name '*.blp' | xgettext --language=C++ --keyword=_ --keyword=L --omit-header -o myapp-gnome-gomod.pot --files-from=-"
//go:generate sh -c "find . -name '*.po' -print0 | xargs -0 -I {} msgmerge --update --backup=none \"{}\" myapp-gnome-gomod.pot"
//go:generate sh -c "find . -type f -name '*.po' -print0 | xargs -0 -I {} sh -c 'mkdir -p $(basename {} .po)/LC_MESSAGES && msgfmt -o $(basename {} .po)/LC_MESSAGES/myapp-gnome-gomod.mo {}'"
//go:generate sh -c "find . -name \"*.po\" -exec basename {} .po \\; > LINGUAS"
//go:embed *
var FS embed.FS
