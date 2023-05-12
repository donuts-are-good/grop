![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

# grop

it's a tool for searching files and their contents

## examples

### searching in a directory

imagine you want to find all instances of the word "error" in log files within a directory `/var/logs`.

#### command
```bash
./grop "error" /var/logs
```
#### output:

```bash
/var/logs/app.log: line 32: "there was an error connecting"
/var/logs/app.log: line 45: "critical error: could not connect"
/var/logs/system.log: line 12: "unexpected error at line 12"
```

### searching in standard input

consider you want to find all instances of the word "john" from another program's output.

#### command:

```
echo -e "glazed\nchocolate\nfilled" | ./grop "glazed"
```

#### output:

```bash
line 1: "glazed"
```

### searching in a file

to search the contents of a specific file, you can use a combination of `cat` and `grop`. for instance, to search for "donut" in `myfile.txt`:

#### command:

```
cat myfile.txt | ./grop "donut"
```

#### output:

```
line 10: "this is some donut in the middle"
line 21: "donut is found again here"
```

## license

MIT License 2023 donuts-are-good, for more info see license.md
