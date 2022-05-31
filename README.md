**&#x1F4CC; This branch was made specifically to scrape the dictionary for ideoms and proverbs &#x1F4CC;**

<br>

# KLED-Scraper

## Motivation

The abbreviation KLED stands for "Korean Learner's English Dictionary" which is based on [Korean Leaner's Dictionary](https://krdict.korean.go.kr/mainAction). The dictionary even provides an extensive API which can be used after registering account on their website.

I found two particular problems with their API that bothered me a lot though:
1. For some reason the requested multimedia part (videos, pronounciation files etc.) is not delivered.
2. The searched keyword is not highlighted in any way in the example sentences.[^1]

For one of my other [projects](https://github.com/Mxngls/kled-server) I wrote an API endpoint that fetches and parses the html for a given word. For future use I thought it might be greate to just scrape the whole (english-korean) dictionary. 

## Description

The given script does exactly that. As the dictionary itself is licensed under the Creative Common License there are no copyright issues to worry about at all.

For example output see the ```dict``` directory. I obtained the necessary IDs by sending an empty prompt while using the advanced search function. The downloadable result (```XML```) contains all entries in there are in 'alphabetical' order.[^2]

## Installation & Usage

Given that Go is installed:
```zsh
% go version go1.17.6 darwin/amd64
```

Just clone the repository:
```zsh
% git clone git@github.com:Mxngls/kled-server.git
````

Then run:
```zsh
% go run .
```

The result will be a JSON file that's roughly 70 MB big and contains all of the more than 52.000 english-korean entries the dictionary has to offer.

[^1]:To see why this might be a problem see the various possible conjugations of the *regular* verb [건네다](https://en.wiktionary.org/wiki/%EA%B1%B4%EB%84%A4%EB%8B%A4#Conjugation).
[^2]:It should be noted that it is questionable if that behavior is intentioned though. .
