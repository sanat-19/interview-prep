# Week 1 - Monday (26 Oct): Arrays + Hashing

Order: easy -> medium. Har problem: brute force -> waste dhoondo -> optimize -> edge cases -> complexity.

## Session 1 (8-11 am): Pattern seekhna
- Hash map kya karta hai: "pehle dekha hai kya?" ka jawab O(1) mein
- Frequency counting pattern
- Guided: Two Sum

## Session 2 (6/7-8 pm): Solve
- Valid Anagram
- First Unique Character in a String
- Group Anagrams (agar time bacha)

## Kaise chalana hai
`go test ./dsa/week1/mon_arrays_hashing/two_sum/` (repo root se). Sab ek saath: `go test ./dsa/week1/mon_arrays_hashing/...`
Shuru mein tests FAIL honge (stubs khaali hain). Tumhara kaam: tests green karna.

## Session 3 (11 pm-12 am): Review
- Tracker mein 3 lines: pattern, key insight, mistake
- Error journal entry

## Tracker
| Problem | Brute force TC | Optimized TC/SC | Stuck? | Re-solve 1d / 1w / 1m |
|---|---|---|---|---|
| Two Sum | | | | |
| Valid Anagram | | | | |
| First Unique Char | | | | |

## Test naming
- `happy: ...` normal cases (brute force pe pehle yeh green karo)
- `edge: ...` edge cases (brute force pe yeh agla step)
- `TestTwoSumTimeLimit` speed check (waste dhoondho -> optimize ke baad green hona chahiye)

Sirf ek group chalane ke liye (repo root se):
`go test -v -run '^TestTwoSum$/happy' ./dsa/week1/mon_arrays_hashing/two_sum/`
`go test -v -run '^TestTwoSum$/edge' ./dsa/week1/mon_arrays_hashing/two_sum/`
