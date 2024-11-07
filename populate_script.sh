# reset database
./gator reset

# register users
./gator register brent
./gator register luna

#login user
./gator login brent

# addfeeds that user wants to follow
./gator addfeed TechCrunch https://techcrunch.com/feed/
./gator addfeed Boot.dev https://blog.boot.dev/index.xml
./gator addfeed IGN https://feeds.feedburner.com/ign/game-reviews

# login different user
./gator login luna


# Have user follow feeds that are already in database
./gator follow https://blog.boot.dev/index.xml


# Have user add a new feed
./gator addfeed HackerNews https://news.ycombinator.com/rss

# Aggerate post from RSS feeds
./gator agg 1m
