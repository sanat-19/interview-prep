# Hints: problem mein ye dikhe to ye try karo

Roz ke hints yahan jodte jaayenge. Naya problem aaye to pehle is table mein dekho ki koi clue match karta hai kya.

## Pattern pehchaan

| # | Problem mein ye dikhe | Ye try karo | Kyun | Example |
|---|---|---|---|---|
| 1 | Baar-baar kuch **search / "pehle dekha hai kya?"** karna pad raha hai | Hash map ya set | Lookup O(1) hota hai, list scan O(n) | Two Sum |
| 2 | Do elements ka **sum / difference target** ho | `need = target - x`, map mein `need` dhoondho | Do loops ki jagah ek loop | Two Sum |
| 3 | **Kitni baar aaya** / count / frequency | Frequency counting: map ya `[26]int` | Ginti ek baar nikalo, baar-baar scan mat karo | Valid Anagram, First Unique Char |
| 4 | Keys ka **chhota fixed set** (jaise a-z) | `[26]int` array, dabba number = `ch - 'a'` | Map se tez, hash ka kaam nahi | Valid Anagram |
| 5 | **Same letters, alag order** (anagram) | Sorted word ya `[26]int` ginti ko **key** banao | Anagrams ki pehchaan ek hi hoti hai | Group Anagrams |
| 6 | Cheezon ko **group / bucket** karna | `map[key][]value`, key = group ki pehchaan | Dibba dhoondhne ka loop hi nahi chahiye | Group Anagrams |
| 7 | **Pehla / first** element jo condition satisfy kare | Pehle pass mein info nikalo, doosre pass mein **shuru se** dekho | Pehle wale ka order bacha rehta hai | First Unique Char |
| 8 | Nested loop mein andar wala loop **wahi cheez baar-baar** dhoondh raha hai | Us cheez ko pehle ek baar nikal ke yaad rakho (map / array) | Yehi "waste" hai, step 4 ka sawaal | First Unique, Group Anagrams |
| 8b | **Top k / sabse zyada / sabse bade k** chahiye | Ginti nikalo, phir sort karo (O(d log d)). Tez chahiye to ginti ko index bana ke buckets | Har baar max scan karna waste hai | Top K Frequent |
| 8c | **Baki sab ka product / sum, khud ko chhodke** (division mana) | Left ka alag, right ka alag banao. `left[i] = left[i-1] * nums[i-1]` | Har `i` pe shuru se nikalna waste hai | Product Except Self |
| 8d | **Sabse lambi consecutive chain** (x, x+1, x+2...) | Set mein rakho. Sirf us number se chain gino jiske `x-1` set mein nahi hai | Beech ke numbers se shuru karna waste hai | Longest Consecutive Sequence |
| 8e | **Baar-baar kisi range ka sum** poocha jaye (ya subarray sum) | Prefix sum ek baar banao. Range = do prefix ka ghatav | Har query mein wahi numbers dobara jodna waste hai | Range Sum Query |
| 8f | **0/1 array, barabar ginti** ya "balanced" hissa | 0 ko -1 maano, prefix sum. Same prefix do jagah = balanced hissa. Map: sum -> pehla index | Har (i,j) pe ginti dobara | Contiguous Array |
| 8g | **Kitne subarrays ka sum k** (negative bhi ho sakte) | Prefix sum + map. `sum - k` pehle kitni baar aaya. Map: sum -> ginti, `{0:1}` se shuru | Har (i,j) ka sum nikalna | Subarray Sum Equals K |
| 9a | **Sorted array**, do numbers ka sum / pair dhoondhna | Do pointers dono ends se. Sum chhota left++, bada right-- | Har jodi (i, j) try karna | Two Sum II |
| 9b | String/array **symmetric** check (palindrome) | left aur right andar ki taraf, non-alphanumeric skip | Naya reversed string banana | Valid Palindrome |
| 9c | Sorted array mein **in place** duplicates hatana | Read aur write pointer, nums[w] = nums[r]. Go mein pop/append nahi, upar likho | append se delete (O(n^2)) | Remove Duplicates |
| 9d | Sirf **0/1/2** (teen values) ko sort | Ginti ya Dutch flag (lo, mid, hi). mid ki value dekho, hi se compare nahi | Poora sort lagana | Sort Colors |
| 9e | Do deewaron/lines ke beech **area** ka max | Ends se two pointers, chhoti deewar wala pointer hilao | Har jodi (i, j) | Container With Most Water |
| 9f | **Triplet / teen numbers** ka sum target, unique results | Sort, ek number fix, baaki do ke liye two pointers. Duplicates i par aur hit ke baad skip | Teen nested loops O(n^3) | 3Sum |
| 9g | **Paani / water between bars**, har index par min(leftMax, rightMax) | Prefix/suffix max arrays, ya sabse unchi deewar tak do loops, ya two pointers | Har index ke liye max dobara nikalna | Trapping Rain Water |
| 14d | Set/map bana hi diya hai, ab uspar loop chahiye | Original array par nahi, `for x := range seen` (duplicates hat jaate hain) | Duplicates se same kaam baar-baar hota hai | Longest Consecutive Sequence |
| 14e | Nested loop mein "is number ko skip karo" | `continue` sabse andar wale loop par asar karta hai. Check bahar wale loop mein, andar se pehle lagao | Warna infinite loop | Longest Consecutive Sequence |
| 14f | Strings ki list ko ek string mein pack karna (kuch bhi character ho sakta hai) | Har string ke aage uski lambai likho (5#hello). Decode lambai padh ke utne bytes uthata hai | Delimiter ya base64 ki zarurat nahi, koi collision nahi | Encode/Decode Strings |
| 14g | Loop mein string jodna (s += x) | strings.Builder use karo | += har baar poori string copy karta hai, O(n^2) | Encode/Decode Strings |

## Galti se bachne ke hints

| # | Situation | Yaad rakho | Example |
|---|---|---|---|
| 9 | Array index letter se banta hai | **Har jagah** same formula: `count[s[i]-'a']` | First Unique (`- 'a'` bhoola) |
| 10 | Current element khud ko match na kare | **Pehle lookup, phir insert** | Two Sum |
| 11 | "Kuch nahi mila" ka faisla | Saare options dekhne ke **baad** (loop ke bahar, `found` flag) | Group Anagrams brute force |
| 12 | Do inputs compare ho rahe hain | Pehle **length alag?** check karo (edge #1) | Valid Anagram |
| 13 | Merged loop mein `+1` aur `-1` saath chal rahe hain | Beech mein negative temporary ho sakta hai, early exit mat lagao | Valid Anagram |
| 14 | Unicode / bada character set | `[26]` array nahi chalega, map use karo | Valid Anagram follow-up |
| 14b | Go ka **map random order** mein ghoomta hai | Order chahiye to `(key, value)` slice banao aur sort karo. Map ko sort nahi kar sakte | Top K Frequent |
| 14c | Sabse bada dhoondhna | `if v > max { max, maxKey = v, x }` (compare sahi taraf aur dono update) | Top K Frequent |

## Complexity ke hints

| # | Situation | Kaise bolna hai |
|---|---|---|
| 15 | Loops **ek ke andar** | Multiply: n × n = O(n²) |
| 16 | Loops **ek ke baad ek** | Add: n + n = 2n = O(n), constants hata do |
| 17 | Map ki space | **Kitni alag keys** (distinct), string ki length nahi. a-z ho to O(1) |
| 18 | Strings ka kaam | `n` words aur har word `k` letters ka ho to `k` bhi bolo (jaise O(n × k log k)) |
