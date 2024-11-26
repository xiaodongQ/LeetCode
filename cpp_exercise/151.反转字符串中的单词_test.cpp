/*
 * @lc app=leetcode.cn id=151 lang=cpp
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
class Solution {
public:
/*
    // 方式1：先实现功能，借助语言内置的反转函数，并且不限制辅助空间占用。
    string reverseWords(string s) {
        // 1）字符串拆分成单词数组 2）数组反转 3）数组拼接
        std::istringstream iss(s);
        vector<string> words;
        string word;
        while (iss >> word) {
            words.push_back(word);
        }
        // 反转，直接使用<algorithm>库中的std::reverse
        std::reverse(words.begin(), words.end());
        // 数组重新拼接
        std::ostringstream oss;
        for (auto i=0; i < words.size(); i++) {
            if (i > 0) {
                oss << " ";
            }
            oss << words[i];
        }
        return oss.str();
    }
*/
    // 方式2：借助双端队列实现反转
    string reverseWords(string s) {
        int left = 0;
        int right = s.size() - 1;
        // 去除前面空格，并记录有效起始位置left
        while (left <= right && s[left] == ' ') {
            left++;
        }
        // 去除后面空格，并记录有效终止位置right
        while (left <= right && s[right] == ' ') {
            right--;
        }
        // 中间处理，并检查空格
        deque<string> words;
        string word;
        while (left <= right) {
            char c = s[left];
            // 新单词
            if (!word.empty() && c == ' ') {
                words.push_front(word);
                word = "";
            } else if (c != ' ') { // 此处单独区分' '，而不是仅else
                word += c;
            }
            left++;
        }
        // 最后一个单词
        words.push_front(word);
        
        // 借助string，拼接新字符串
        string result;
        for (auto i = 0; i < words.size(); i++) {
            if (i > 0) {
                result += " ";
            }
            result += words[i];
        }
        return result;
    }
};
// @lc code=end

