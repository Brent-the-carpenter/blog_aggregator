# reset database
./blog_aggregator reset

# register users
./blog_aggregator register brent
./blog_aggregator register luna

#login user
./blog_aggregator login brent

# addfeeds that user wants to follow
./blog_aggregator addfeed TechCrunch https://techcrunch.com/feed/
./blog_aggregator addfeed Boot.dev https://blog.boot.dev/index.xml
./blog_aggregator addfeed IGN https://feeds.feedburner.com/ign/game-reviews

# login different user
./blog_aggregator login luna


# Have user follow feeds that are already in database
./blog_aggregator follow https://blog.boot.dev/index.xml


# Have user add a new feed
./blog_aggregator addfeed HackerNews https://news.ycombinator.com/rss

# Aggerate post from RSS feeds
./blog_aggregator agg 1m
