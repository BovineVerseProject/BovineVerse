package sensitive

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"time"
)

var df *filter

// Filter 敏感词过滤器
type filter struct {
	trie       *Trie
	noise      *regexp.Regexp
	buildVer   int64
	updatedVer int64
}

// InitSensitiveFilter 初始化敏感词库
func InitSensitiveFilter(path string) error {
	df = &filter{
		trie:  newTrie(),
		noise: regexp.MustCompile(`[\|\s&%$@*]+`),
	}

	return df.loadWordDict(path)
}

// UpdateNoisePattern 更新去噪模式
func (filter *filter) UpdateNoisePattern(pattern string) {
	filter.noise = regexp.MustCompile(pattern)
}

// LoadWordDict 加载敏感词字典
func (filter *filter) loadWordDict(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return filter.Load(f)
}

// Load common method to add words
func (filter *filter) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		filter.addWord(string(line))
	}

	return nil
}

func (filter *filter) updateFailureLink() {
	if filter.buildVer != filter.updatedVer {
		// fmt.Println("update failure link")
		filter.trie.BuildFailureLinks()
		filter.buildVer = filter.updatedVer
	}
}

// AddWord 添加敏感词
func AddWord(words ...string) {
	df.addWord(words...)
}

func (filter *filter) addWord(words ...string) {
	filter.trie.Add(words...)
	filter.updatedVer = time.Now().UnixNano()
}

// Filter 过滤敏感词
func Filter(text string) string {
	return df.filter(text)
}

func (filter *filter) filter(text string) string {
	filter.updateFailureLink()
	return filter.trie.Filter(text)
}

// Replace 和谐敏感词
func Replace(text string, repl rune) string {
	return df.replace(text, repl)
}
func (filter *filter) replace(text string, repl rune) string {
	filter.updateFailureLink()
	return filter.trie.Replace(text, repl)
}

// FindIn 检测敏感词
func FindIn(text string) (bool, string) {
	return df.findIn(text)
}
func (filter *filter) findIn(text string) (bool, string) {
	filter.updateFailureLink()
	text = filter.removeNoise(text)
	return filter.trie.FindIn(text)
}

// FindAll 找到所有匹配词
func (filter *filter) FindAll(text string) []string {
	filter.updateFailureLink()
	return filter.trie.FindAll(text)
}

// Validate 检测字符串是否合法
func Validate(text string) (bool, string) {
	return df.validate(text)
}
func (filter *filter) validate(text string) (bool, string) {
	filter.updateFailureLink()
	text = filter.removeNoise(text)
	return filter.trie.Validate(text)
}

// RemoveNoise 去除空格等噪音
func RemoveNoise(text string) string {
	return df.removeNoise(text)
}

func (filter *filter) removeNoise(text string) string {
	return filter.noise.ReplaceAllString(text, "")
}
