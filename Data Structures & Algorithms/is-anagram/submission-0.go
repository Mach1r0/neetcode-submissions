func isAnagram(s string, t string) bool {
  counts := make(map[byte]int, len(s))

  if len(s) != len(t){ 
    return false
  }

  for i := 0; i < len(s); i++ {
      counts[s[i]]++ 
      counts[t[i]]-- 
  }
  
 for _, count := range counts {
    if count != 0 {
        return false 
    }
}

return true 
} 
