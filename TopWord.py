#!/usr/bin/python
# -*- coding:utf-8 -*-


import re
import requests


wordTotalCount = 86800

def wordApi(page):
    response = requests.get("http://www.wordcount.org/dbquery.php?toFind=%d&method=SEARCH_BY_INDEX" % page)
    return response.text

def Top(str,page):
    with open("cache/%s.cache" % page,"w") as f:
        f.write(str)
    words = re.findall(r"word(\d+)=([\w\-\']+)&",str)
    if words:
        with open("words/%s.md" % page,"w") as f:
            for word in words:
                f.write("%s. %s\n" % (word[0],word[1]))
        print "now saved %d word" % (page-1)
    else:
        raise Exception("not found words")

if __name__ == "__main__":
    page = 86400
    while True:
        str = wordApi(page)
        Top(str,page)
        page += 300
        if page >= 86800:
            break


    #Top()
