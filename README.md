# jsonf
A Unix-y utility for formatting JSON in a more stream-friendly way.

# Usage

`jsonf` reads from `stdin` and writes to `stdout` by default.

```bash
$ jsonf --help
Usage of jsonf:
  -a="[]": The runes used to wrap output array indexes.
  -i="": The input file. Defaults to stdin.
  -k="{}": The runes used to wrap output keys.
  -o="": The output file. Defaults to stdout.
  -s="\"\"": The runes used to wrap output strings.
  -sep="->": String used to seperate keys.
```

```bash
$ jsonf
{"colors": ["red", "blue", "green"], "numbers": [1, 2, 3], "active": true}
->{active}
->{active}->true
->{colors}
->{colors}->[]
->{colors}->[0]->"red"
->{colors}->[1]->"blue"
->{colors}->[2]->"green"
->{numbers}
->{numbers}->[]
->{numbers}->[0]->1
->{numbers}->[1]->2
->{numbers}->[2]->3
```

# Examples

### Get information about the most recent commit on this repository

```bash
$ curl -s 'https://api.github.com/repos/andrewguenther/jsonf/commits' |
  jsonf -sep='>' |
  grep -a "\[0\].*{commit}>{committer}>.*>" |
  cut -d ">" -f 5,6
{name}>"Andrew Guenther"
{email}>"guenther.andrew.j@gmail.com"
{date}>"2015-05-02T09:24:43Z"
```

### Get all the front page links on Reddit

```bash
$ curl -s 'https://www.reddit.com/.json' |
  jsonf -sep='>' |
  grep '>{data}>{children}>\[.\+\]>{data}>{url}>' |
  cut -d '>' -f 7
"https://i.imgur.com/qH91IEl.gifv"
"http://imgur.com/MUXn9fD"
"http://www.mnn.com//earth-matters/wilderness-resources/stories/more-trees-than-there-were-100-years-ago-its-true"
"http://imgur.com/mjbuXpA"
"http://i.imgur.com/XJ2SG5J.jpg"
"http://imgur.com/tvUSaLZ"
"http://www.reddit.com/r/science/comments/34kxuh/do_you_have_a_college_degree_or_higher_in_science/"
"http://imgur.com/qOPPuzb"
"http://www.reuters.com/article/2015/04/30/us-usa-police-apps-idUSKBN0NL2SK20150430"
"https://m.youtube.com/watch?v=pNHpDG9dZq0"
"http://www.thefix.com/content/ecuador-considers-decriminalizing-all-illegal-drug-use"
"http://i.imgur.com/9x6acUY.jpg"
"http://www.hollywoodreporter.com/news/adam-sandlers-ridiculous-six-makeup-792582"
"http://www.reddit.com/r/AskReddit/comments/34kkcs/pregnant_people_of_reddit_whats_the_rudestmost/"
"http://imgur.com/5Q1RzVV"
"http://www.reddit.com/r/Jokes/comments/34kihw/how_many_reddit_mods_does_it_take_to_change_a/"
"http://imgur.com/a/8aFdi"
"http://www.cbsnews.com/news/woman-texted-driving-drunk-woo-just-before-fatal-crash/"
"http://imgur.com/a/Rzrk9"
"http://www.reddit.com/r/Showerthoughts/comments/34j0r7/that_google_chrome_what_tab_is_being_noisy/"
"http://i.imgur.com/vuBumaG.jpg"
"http://www.reddit.com/r/WritingPrompts/comments/34j4ph/wp_elon_musk_is_actually_a_stranded_alien_who/"
"http://www.reddit.com/r/IAmA/comments/34k027/i_am_noel_gallagher_amaa/"
"http://imgur.com/O4hkbVt"
"http://imgur.com/a/lDn0S"
```
