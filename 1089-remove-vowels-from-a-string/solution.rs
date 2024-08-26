impl Solution {
    pub fn remove_vowels(s: String) -> String {
        let mut res = String::new();  // Initialize an empty String

        for c in s.chars() {
            if !matches!(c, 'a' | 'e' | 'i' | 'o' | 'u' | 'A' | 'E' | 'I' | 'O' | 'U') {
                res.push(c);  // Push the character to the result string
            }
        }

        res
    }
}
