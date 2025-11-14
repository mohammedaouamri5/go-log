package log

import "github.com/charmbracelet/lipgloss"

func DefaultTheme() *Colors {
	return &Colors{

		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#00C7FF")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true).Italic(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Background(lipgloss.Color("#2E0000")).Bold(true).Underline(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#999999")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#87CEEB")).Bold(true), // skyblue
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4")).Bold(true), // pink

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FFEA00")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#75715E")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E6DB74")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#AE81FF")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FD971F")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#75715E")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#66D9EF")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")),
	}
}

func MonokaiTheme() *Colors {
	return &Colors{
		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#66D9EF")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E6DB74")).Bold(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#FD971F")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#F92672")).Background(lipgloss.Color("#2B1B1D")).Bold(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#75715E")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#A6E22E")).Bold(true),
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#F92672")).Bold(true),
		Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#66D9EF")).Italic(true),

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#F92672")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#75715E")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E6DB74")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#AE81FF")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FD971F")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#75715E")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#66D9EF")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")),
	}
}
func DraculaTheme() *Colors {
	return &Colors{
		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#F1FA8C")).Bold(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB86C")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Background(lipgloss.Color("#282A36")).Bold(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B")).Bold(true),
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6")).Bold(true),
		Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9")).Italic(true),

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#F1FA8C")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#44475A")),
	}
}
func GruvboxTheme() *Colors {
	return &Colors{
		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#83A598")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FABD2F")).Bold(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#FE8019")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#FB4934")).Background(lipgloss.Color("#3C3836")).Bold(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#928374")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#B8BB26")).Bold(true),
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#FB4934")).Bold(true),
		Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#D3869B")).Italic(true),

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#D3869B")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#928374")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#B8BB26")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#FABD2F")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FE8019")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#928374")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#83A598")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#7C6F64")),
	}
}
func SolarizedDarkTheme() *Colors {
	return &Colors{
		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#268BD2")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#B58900")).Bold(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#CB4B16")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#DC322F")).Background(lipgloss.Color("#002B36")).Bold(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#586E75")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#859900")).Bold(true),
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#DC322F")).Bold(true),
		Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#6C71C4")).Italic(true),

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#2AA198")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#586E75")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#B58900")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#6C71C4")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#CB4B16")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#586E75")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#268BD2")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#073642")),
	}
}
func NordTheme() *Colors {
	return &Colors{
		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#88C0D0")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#EBCB8B")).Bold(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#D08770")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#BF616A")).Background(lipgloss.Color("#2E3440")).Bold(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#4C566A")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#A3BE8C")).Bold(true),
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#BF616A")).Bold(true),
		Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#B48EAD")).Italic(true),

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#81A1C1")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#4C566A")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#EBCB8B")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#88C0D0")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#D08770")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#4C566A")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#8FBCBB")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#434C5E")),
	}
}
func OneDarkTheme() *Colors {
	return &Colors{
		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#61AFEF")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E5C07B")).Bold(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#D19A66")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#E06C75")).Background(lipgloss.Color("#282C34")).Bold(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#5C6370")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#98C379")).Bold(true),
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#E06C75")).Bold(true),
		Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#C678DD")).Italic(true),

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#56B6C2")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#5C6370")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E5C07B")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#61AFEF")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#D19A66")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#5C6370")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#61AFEF")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#3E4451")),
	}
}
func TokyoNightTheme() *Colors {
	return &Colors{
		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E0AF68")).Bold(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#F7768E")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#F7768E")).Background(lipgloss.Color("#1A1B26")).Bold(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#565F89")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#9ECE6A")).Bold(true),
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#F7768E")).Bold(true),
		Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#BB9AF7")).Italic(true),

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#7DCFFF")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#565F89")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E0AF68")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#F7768E")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#565F89")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#7DCFFF")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#3B4261")),
	}
}
func AyuDarkTheme() *Colors {
	return &Colors{
		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#39BAE6")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E6B450")).Bold(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF7733")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF3333")).Background(lipgloss.Color("#0F1419")).Bold(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#5C6773")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#AAD94C")).Bold(true),
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF3333")).Bold(true),
		Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB454")).Italic(true),

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#36A3D9")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#5C6773")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E6B450")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#39BAE6")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF7733")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#5C6773")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#39BAE6")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#3E4B59")),
	}
}
func LightTheme() *Colors {
	return &Colors{
		Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#005F87")).Bold(true),
		Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#AF8700")).Bold(true),
		Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#D75F00")).Bold(true),
		Fatal: lipgloss.NewStyle().Foreground(lipgloss.Color("#AF0000")).Background(lipgloss.Color("#FFFFFF")).Bold(true),

		Time: lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Italic(true),
		File: lipgloss.NewStyle().Foreground(lipgloss.Color("#5FAF00")).Bold(true),
		Line: lipgloss.NewStyle().Foreground(lipgloss.Color("#AF0000")).Bold(true),
		Func: lipgloss.NewStyle().Foreground(lipgloss.Color("#875FAF")).Italic(true),

		KeyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#005F87")).Bold(true),
		TypeStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Italic(true),
		StrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#AF8700")),
		NumStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#5F5FAF")),
		BoolStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#D75F00")).Bold(true),
		NullStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Italic(true),
		PtrStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#005F87")).Italic(true),
		Bracket:   lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA")),
	}
}
