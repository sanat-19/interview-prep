# DSA notes: 12 hafte ka learning log

Ye file roz update hoti hai. 12 hafte baad padh ke samajh aana chahiye ki kab kya seekha, kaunsi galti kari, aur kaunsa pattern kab lagta hai.

## Har problem ka flow
1. Question + real-life example + animation
2. Happy path brute force (khud likho)
3. Edge cases brute force pe (checklist: `EDGE_CASES.md`; pehle khud likho, phir tests)
4. Waste dhoondho
5. Optimize

Hints ka table: `HINTS.md`. Edge case checklist: `EDGE_CASES.md`.

Complexity ki quick yaad: time = kaam kitna badhta hai input ke saath. Space = extra jagah kitni lagi (input ki jagah nahi ginte). Map ki space = **kitni alag keys**, string ki length nahi.

---

## Week 1, Monday 26 Oct: Arrays + Hashing

Pattern: **hash map = "pehle dekha hai kya?" ka jawab O(1) mein**. Do roop: (a) complement dhoondhna, (b) frequency counting.

### Two Sum (easy)
- Sawaal: array mein do numbers jinka sum target ho, unke indices.
- Brute force: do loops, har pair check. Time O(n^2), space O(1).
- Waste: har number ke liye poori list dobara scan ho rahi hai.
- Optimized: map `number -> index`. Har number pe `need = target - n` map mein dhoondo, mila to answer, nahi to current number map mein daalo. Time O(n), space O(n).
- Zaroori detail: **pehle lookup, phir insert.** Isse current number khud ko nahi dhoondh paata (`[3,2,4]`, target 6 mein 3 apna hi jodidaar na ban jaye).
- Galti: pehle map mein key aur value ulti soch li thi (index->number). Lookup number se hota hai, isliye key number hai.
- Do-pass (pehle sab map mein, phir lookup) bhi chalta hai: O(2n) matlab O(n). Constants hata dete hain. Par apne aap ko dobara dhoondhne ki trap handle karni padti hai.
- Interview mein bolna: "Brute force O(n^2) hai. Hash map se complement dhoondh ke O(n) time, O(n) space."

### Valid Anagram (easy)
- Sawaal: kya `t`, `s` ke letters ko aage-peeche karke ban sakta hai (same letters, same count)?
- Real-life: Scrabble ki do thaili, tiles ek-ek karke cut karo.
- Pattern: **frequency counting**. Ek map mein `s` ke letters ginti karo, `t` ke letters se ghatao.
- Complexity: time O(n). Space **O(1)**, kyunki map mein max 26 keys (a-z), chahe string kitni bhi lambi ho. Unicode ho to O(k), k = alag characters.
- Edge case jo miss hua: **`len(s) != len(t)`** (checklist #1, size). Code `s` ko hi bada maan ke chal raha tha, `t` ke extra letters ignore ho rahe the. Fix: sabse upar ek line, length alag hai to seedha `false`.
- Sabak: pehle sawaal ko saaf karo ("alag length ke anagram ho sakte hain?"), phir code. Mera bada/chhota dhoondhne wala plan zyada ghuma hua tha.
- Step 4 (waste dhoondhna): map hataya, uski jagah **`[26]int` array**. Keys fixed hain (a-z), to dabba number seedha `letter - 'a'` se milta hai: `'a'-'a'=0`, `'c'-'a'=2`, `'z'-'a'=25`. Hash nikalne ka kaam bachta hai. Big-O wahi (time O(n), space O(1)), par constant chhota aur code tez.
- Loops 3 se 2 kiye: lengths barabar hain (pehle check ho chuka), to ek hi loop mein `count[s[i]-'a']++` aur `count[t[i]-'a']--` chal sakta hai. Phir 26 dabbe dekho: sab 0 hain to anagram, koi bhi non-zero hai to nahi.
- Final shape:
  ```go
  if len(s) != len(t) { return false }
  count := [26]int{}
  for i := 0; i < len(s); i++ { count[s[i]-'a']++; count[t[i]-'a']-- }
  for i := 0; i < len(count); i++ { if count[i] != 0 { return false } }
  return true
  ```
- Dhyan rakho: array trick sirf tab chalti hai jab characters ka chhota fixed set ho (a-z). Unicode ho to map wapas.
- Interview mein bolna: "Dono strings ki length alag ho to false. Warna 26 size ka count array, s ke liye +1, t ke liye -1, end mein sab 0 hone chahiye. O(n) time, O(1) space."
- Abhi baaki: tile-cut brute force ki complexity (O(n^2) kyun) aur 3-4 aur apne edge cases khud likhna.

### First Unique Character in a String (easy)
- Sawaal: string mein pehla letter jo sirf ek baar aaya ho, uska **index** (nahi mila to -1).
- Real-life: line mein khade players, pehla jiski jersey number koi aur use nahi kar raha.
- Brute force: bahar loop har index `i`, andar loop poori string mein `s[i]` kitni baar aaya. count == 1 to `i` return. Time O(n^2), space O(1).
- Edge cases (sab brute force pe khud pass): `""` -> -1 (#1), `"z"` -> 0 (#1), `"aaaa"` -> -1 (#5), `"aabb"` -> -1 (#7), `"aabbc"` -> 4 (#6), `"ab"` -> 0.
- Waste: ek hi letter ('e' jo 4 baar aaya) ke liye poori string baar-baar scan ho rahi thi, wahi ginti dobara dobara nikal rahe the. n = 1 lakh pe ~10^10 steps.
- Optimized (frequency counting, 2 alag loops): pehle loop mein `count[s[i]-'a']++` se saari ginti ek baar. Doosre loop mein string shuru se dekho, jis letter ka `count[s[i]-'a'] == 1` ho uska index return. Time O(n) (n + n, 2n bhi O(n)), space O(1) (26 dabbe).
- Zaroori: **do loops ek ke andar nahi, ek ke baad ek**, isliye n^2 nahi, n + n.
- Galti: doosre loop mein `count[s[i]]` likha, `- 'a'` bhool gaye -> index out of range ('l' ka number 108, array mein sirf 0..25). Dono jagah wahi formula chahiye.
- Interview mein bolna: "Pehle pass mein 26 size ke array mein frequency, doosre pass mein shuru se pehla letter jiski frequency 1 ho. O(n) time, O(1) space."

### Group Anagrams (medium)
- Sawaal: strings ki list mein jo anagram hain unhe ek group mein daalo (group aur andar ka order kuch bhi).
- Real-life: party mein har table pe wahi log jinke tile-bag ke letters same hon.
- Brute force: har word ke liye existing dibbon ke pehle word se `isAnagram` check; mila to usi dibbe mein, nahi to naya dibba. Time O(n^2 * k), space O(n * k).
- Galti: "koi nahi mila, naya dibba banao" ka `else` dibbon ke loop **ke andar** likha tha. Ye faisla saare dibbe dekhne ke **baad** hota hai (`found` flag + loop ke baad `if !found`). Saath mein shuru mein `group` khaali tha to loop chalta hi nahi tha, kuch group nahi bana.
- Waste: har word ke liye saare dibbe aur har baar `isAnagram`. Dibba dhoondhne ka loop hi bekaar hai.
- Optimized: har group ki **pehchaan (key)** banao. Word ke letters **sort** kar do ("eat", "tea", "ate" -> sab "aet"). Phir `map[string][]string` mein `groups[key] = append(groups[key], word)`. Map ki values = answer.
- Complexity: time O(n * k log k) (har word ki sorting), space O(n * k).
- Aage ka idea: sort ke bajaye `[26]int` ginti ko hi key bana sakte ho (Go mein array map ki key ban sakta hai), to sort ka log k hat jata hai: O(n * k).
- Interview mein bolna: "Anagrams ki sorted spelling same hoti hai, to sorted word ko hash map ki key banaya. Ek pass mein group, O(n * k log k)."

## Week 1, Monday raat: Top K

### Top K Frequent Elements (medium)
- Sawaal: array `nums` aur `k`. Sabse zyada baar aane wale `k` numbers wapas karo (order kuch bhi).
- Real-life: class monitor ka chunav. Parchiyan gino, sabse zyada vote wale k jeet gaye.
- Brute force: pehle `map[int]int` mein ginti. Phir `k` baar: map mein sabse bada dhoondho, result mein daalo, `delete` karo. Time O(n + k * d), space O(d) (d = alag numbers).
- Galtiyan: (1) `k` ka use hi nahi kiya, (2) Go ka map **random order** mein ghoomta hai, to "pichle se bada" compare bekaar hai, (3) `if max > v` ulta likha aur `max` update nahi kiya. Sahi: `if v > max { max, maxKey = v, x }`.
- Space: O(1) bola tha, galat. Numbers ka fixed chhota set nahi hai (a-z jaisa), to map O(d) = O(n).
- Waste: har round mein poora map dobara scan.
- Optimize 1 (sort): map ke `(num, cnt)` jode slice mein daalo, `sort.Slice` ginti ke hisaab se descending, pehle `k`. **Map ka order yaad nahi rehta**, isliye map ko sort nahi kar sakte, slice mein badalna padta hai. Time O(n + d log d).
- Optimize 2 (buckets): `buckets[c]` = wo numbers jinki ginti `c` hai (ginti 1 se n ke beech hi hoti hai). Upar se neeche aate hue k utha lo. Time O(n), space O(n). Buckets complex lagte hain, interview mein pehle sort wala bolo, buckets follow-up mein.
- Interview mein bolna: "Ginti ke liye hash map, phir ginti ke hisaab se sort karke top k. O(n + d log d). Chaho to bucket sort se O(n)."

### Product of Array Except Self (medium)
- Sawaal: `result[i]` = `nums[i]` ko chhodke baaki sabka product. Division mana hai.
- Brute force: har `i` ke liye andar ka loop, `i != j` wale guna karo. Time O(n^2), space O(1) (result ko nahi ginte). Time-limit test fail hota hai.
- Waste: har `i` ke liye left wale aur right wale ka product shuru se dobara nikal rahe the.
- Idea (prefix / suffix): `left[i]` = i ke left ka product, `left[0] = 1`, `left[i] = left[i-1] * nums[i-1]`. `right[i]` = i ke right ka product, `right[n-1] = 1`, `right[i] = right[i+1] * nums[i+1]`. `result[i] = left[i] * right[i]`.
- Time O(n) (3 alag loop jodte hain), space O(n) (`left` aur `right`). Maine pehle O(1) bola tha, galat.
- Bonus (O(1) space): `left` result mein hi daalo, phir peeche se ek variable `r` rakho. Pehle `result[i] *= r`, phir `r *= nums[i]`. Order ulta hua to galat.
- `if i == 0` loop ke andar nahi, loop se bahar likho.
- Interview mein bolna: "left aur right ke product alag banata hoon, dono ka guna. O(n) time, O(n) space. Right ke liye variable rakh ke O(1) space."

## Week 1, Tuesday subah: hashing hard

### Longest Consecutive Sequence (medium)
- Sawaal: array mein se sabse lambi chain (x, x+1, x+2...) ki lambai. Numbers array mein paas paas hone zaroori nahi. Target O(n).
- Real-life: bikhri hui roll number ki parchiyan, sabse lambi lagatar line.
- Brute force: set (map) banao. Har number se shuru karke poocho "x+1 hai kya?" jab tak jawab nahi. O(n^2) time, O(n) space.
- Galtiyan: (1) `if` se sirf ek baar poochha, chain ke liye `for` chahiye. (2) `max` ka check `if` ke andar tha, to akele number ka jawab 0 aaya. (3) `continue` andar wale loop mein tha, infinite loop. (4) `len(seen)` tak chalaya par `nums[i]` uthaya, `[0,0,-1]` par -1 miss.
- Waste: (1) `j` loop hamesha n baar chalta hai. (2) chain ke beech ke numbers se bhi shuru karte hain, wahi cheezein dobara ginte hain.
- Optimize: sirf chain ki shuruat se gino, yaani jiske `x-1` seen mein nahi. `for seen[x+count]` se chain badhao. Har number ek hi chain mein ek baar, O(n) time, O(n) space.
- Sort wala tareeka bhi sahi hai: sort karke ek pass, duplicates skip. O(n log n).
- Duplicate trap: `nums` par loop chalane se duplicate shuruat wale number baar-baar poori chain chalate hain (18 second vs 4 ms). `seen` ki keys par loop chalao.
- Rule: "beech wale ko mat chhedo, sirf shuruat se gino."
- Interview mein bolna: "Brute force O(n^2). Sort karke O(n log n). Set mein rakh ke sirf chain ki shuruat se gino: O(n)."

### Encode and Decode Strings (medium)
- Sawaal: list of strings ko ek string mein pack karo (Encode), aur wapas wahi list banao (Decode). Strings mein kuch bhi ho sakta hai, khaali bhi.
- Pehla tareeka: har string alag base64, "," se jodo. Base64 mein comma hota hi nahi, to split safe. (Poori concatenation ka base64 nahi chalta, boundaries gum ho jaati hain.)
- Galtiyan: (1) Encode mein URLEncoding, Decode mein StdEncoding, alag alphabet. Error ignore kiya, to unicode test fail. (2) khaali list aur ek khaali string dono ka encode "", collision. Fix: har item ke baad "," ya length-prefix.
- Final (length-prefix): har string ke aage lambai#. ["hello","","a#b"] -> 5#hello0#3#a#b. Decode lambai padhta hai aur utne bytes uthata hai, andar kuch bhi ho. O(total length). strings.Builder use karo, s += x baar-baar copy karta hai (O(n^2)).
- Seekh: error ko _ se ignore mat karo. Length-prefix ek common pattern hai (HTTP Content-Length, Redis RESP).

## Week 1, Tuesday: prefix sum

### Range Sum Query - Immutable (easy)
- Sawaal: ek array diya, phir baar-baar SumRange(left, right) poochha jaata hai. Constructor ek baar chalta hai, SumRange baar-baar.
- Yahan input/output ek function ka nahi, ek object ka hai: Constructor(nums) se object banta hai, phir har SumRange(left, right) ek number deta hai. left aur right bahar se aate hain.
- Real-life: charger ka meter ek running total hai. meterStop - meterStart = energy. Wahi prefix sum hai, charger ne pehle se kar rakha hai.
- Brute force: har SumRange mein left se right tak loop. O(n) per query, O(n*q) total, time-limit fail.
- Waste: har query mein wahi numbers dobara jodte hain.
- Optimize (prefix sum): Constructor mein ek baar pre banao. SumRange O(1). Constructor O(n), space O(n).
- Do tareeke: (1) pre[i] = nums[0..i] ka sum, bina extra 0. Tab left == 0 ke liye alag case: pre[right], warna pre[right] - pre[left-1]. (2) shuru mein extra 0: pre[i+1] = pre[i] + nums[i], SumRange = pre[right+1] - pre[left]. Koi special case nahi.
- Tumhare version mein make([]int, len(nums)+1) tha, ek element extra hai, bekaar par nuksaan nahi.
- Interview mein bolna: "Ek baar prefix sum banata hoon, phir har range ka sum do prefix ka ghatav, O(1)."

### Contiguous Array (medium)
- Sawaal: 0/1 array mein sabse lamba continuous hissa jisme 0 aur 1 barabar hon. Jawab = length.
- Brute force: har start i, har end j, zeros/ones gino. O(n^2) time, O(1) space. Time test fail.
- Waste: har i par wahi hisse dobara dekhte hain.
- Trick: 0 ko -1 aur 1 ko +1 maano. Balanced hissa ka sum 0. pre[j+1] == pre[i] matlab beech ka hissa balanced.
- Optimize: map[sum] = PEHLA index. Map {0: -1} se shuru. Same sum dobara aaye to length = i - pehla index. Sirf naya sum aaye to store, taaki pehla index bana rahe. O(n) time, O(n) space.
- Do version: ek pass (running sum + map), ya teen pass (pre array, seen map, max). Dono sahi, kyunki yahan pehla index store hota hai jo hamesha i se pehle ya barabar hota hai.

### Subarray Sum Equals K (medium)
- Sawaal: kitne continuous hisson ka sum exactly k. Length nahi, ginti. Numbers negative ya 0 bhi ho sakte hain, to sort ya window nahi chalega.
- Brute force: har (i, j), sum k ho to count++. O(n^2).
- Idea (charger meter): sum - purani_reading = k, to purani_reading = sum - k. Map se pucho ye pehle kitni baar aayi.
- Map: map[sum] = KITNI BAAR aaya (ginti, index nahi), kyunki har baar ek alag group banta hai.
- Lookup pehle, add baad mein, ek hi loop mein. Alag loop mein seen poora bana do to aage wale prefix bhi gin jaate hain, galat. Tests fail hue the.
- Do tareeke start-case ke: map {0: 1} se shuru, ya khaali map aur alag `if pre[i] == k {count++}`. Dono sahi.
- O(n) time, O(n) space.

---

## Mistake log (weak spots dekhne ke liye)
| Date | Problem | Mistake | Checklist # |
|---|---|---|---|
| 26 Oct | Two Sum | map ki key/value ulti | - |
| 26 Oct | Valid Anagram | length alag wala case miss | 1 |
| 26 Oct | Valid Anagram | space complexity samajh nahi aayi (O(n) bola, O(1) hai) | - |
| 26 Oct | First Unique Char | `count[s[i]]` mein `- 'a'` bhool gaye (index out of range) | - |
| 26 Oct | Group Anagrams | naya dibba banane ka `else` dibbon ke loop ke andar likha (faisla loop ke baad hota hai) | 1 |

---

## Schedule aur catch-up (Week 1)
- Timing ab: subah 9-11 am, raat 11 pm-1 am, roz.
- Thu 29 Oct se Sun 1 Nov: LASIK recovery ke liye padhai band. Doctor ki instructions pehle.
- Mon-Wed ka plan: Mon = hashing (4 problems done) + raat Top K Frequent, Product of Array Except Self. Tue = hashing hard (Longest Consecutive Sequence, Encode/Decode Strings) + raat Two Pointers. Wed = HLD URL Shortener + raat LLD Parking Lot.
- Jo miss hoga (Week 1 ka Sat mixed set, Sun revision, referral list, resume draft, behavioral stories): **Sat 7 Nov - Sun 8 Nov** mein cover. Week 2 apna kaam Mon 2 Nov se normal chalayega (agar aankhein theek hon).
- Catch-up weekend: Sat = Week 1 mixed set (pehla aadha) + Week 2 mixed set (doosra aadha). Sun = dono hafton ka revision + error journal + referral list.
| 26 Oct | Top K Frequent | `k` ignore, map ka random order, `v > max` ulta + `max` update nahi | - |
| 26 Oct | Product Except Self | space O(1) bola, par `left` aur `right` do array the, O(n) hai | - |

| 27 Oct | Longest Consecutive | `if` ki jagah `for`, max ka check `if` ke andar, `continue` andar wale loop mein, `len(seen)` tak chala ke `nums[i]` uthaya | - |

| 27 Oct | Encode/Decode Strings | URLEncoding vs StdEncoding mismatch, error ignore kiya, khaali list aur [""] ka collision | - |

| 27 Oct | Range Sum Query | brute force O(n*q); `left == 0` ke bina `pre[-1]` ka khatra | - |

| 27 Oct | Contiguous Array | brute force mein `if zeros == ones` inner loop ke bahar; map key/value ulte; `sum` update bhool gaye; lookup `seen[nums[i]]` ki jagah `seen[sum]` | `{0: -1}` aur pehla index hi kyun |
| 27 Oct | Subarray Sum Equals K | alag loop mein seen bharna (aage wale prefix bhi gin jaate hain) | lookup pehle, add baad mein; map mein ginti |

## Naya weekly schedule (26 Oct se)
- Mon: naya pattern (easy-medium). Tue: usi pattern ke hard. Wed: doosra pattern (easy-medium). Thu: usi ke hard.
- Fri: LLD (half day). Busy ho to Sat. Sat: HLD (half day). LLD Sat par gaya to HLD Sun.
- Sun: flexible. Mixed session, failed problems ki revision, catch-up. Baaki din off.
- Week 1: HLD (URL Shortener) aur LLD (Parking Lot) drop, LASIK ki wajah se. Thu 29 Oct se Sun 1 Nov skipped.
- Wed 28 Oct: Two Pointers (easy-medium). Hard wale (3Sum, Container, Trapping Rain Water), resume draft, stories, referral list: Sun 8 Nov.
- Rule: koi bhi din zero nahi. Kam se kam 1 problem. Do din lagatar miss hue to plan halka karenge.
