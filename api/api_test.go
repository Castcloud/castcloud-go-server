package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

var testServer *httptest.Server
var testRSS string
var testAtom string

func TestMain(m *testing.M) {
	storePath := tempfile()
	defer os.Remove(storePath)

	openStore(storePath)
	authCache = newMemAuthCache()
	initTestData()
	defer store.Close()

	testServer = httptest.NewServer(http.HandlerFunc(testHandler))
	defer testServer.Close()
	testRSS = testServer.URL + "/rss"
	testAtom = testServer.URL + "/atom"

	crawl = newCrawler(time.Hour)
	crawl.start(4)

	os.Exit(m.Run())
}

func initTestData() {
	store.AddUser(&User{
		Username: "test",
		Password: "pass",
	})
	store.AddClient(1, &Client{
		Token: "token",
		UUID:  "real_unique",
		Name:  "Castcloud",
	})
	store.SaveCast(&Cast{
		URL:  "test.go",
		Name: "test",
	})
	store.SaveEpisode(&Episode{
		CastID: 1,
		GUID:   "guid",
	})
	store.SaveEpisode(&Episode{
		CastID:  69,
		GUID:    "since1",
		CrawlTS: 32503680000,
	})
	store.SaveEpisode(&Episode{
		CastID:  69,
		GUID:    "since2",
		CrawlTS: 32503680001,
	})
	store.AddSubscription(1, 1)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	switch r.URL.Path {
	case "/rss":
		w.Write(testFeed)

	case "/atom":
		w.Write(atomTestFeed)

	case "/notxml":
		w.Write([]byte("What is this stuff?"))

	case "/notfeed":
		w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
			<cake><pie taste="so_gewd"></pie></cake>`))

	case "/badrss":
		w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
			<rss><channel></channel></rss>`))

	case "/badatom":
		w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
			<feed></feed>`))

	default:
		w.WriteHeader(404)
	}
}

type testReq struct {
	h http.Handler
	*http.Request
}

func testRequest(h http.Handler, method, url string, body io.Reader) testReq {
	req, _ := http.NewRequest(method, url, body)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return testReq{h: h, Request: req}
}

func (t testReq) send() *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	t.h.ServeHTTP(w, t.Request)
	return w
}

func testForm(f url.Values) *bytes.Buffer {
	return bytes.NewBufferString(f.Encode())
}

func testJSON(v interface{}) string {
	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(v)
	return buf.String()
}

func tempfile() string {
	f, _ := ioutil.TempFile("", "castcloud_")
	f.Close()
	return f.Name()
}

var testFeed = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" media="screen" href="/~d/styles/rss2enclosuresfull.xsl"?><?xml-stylesheet type="text/css" media="screen" href="http://feeds.feedburner.com/~d/styles/itemcontent.css"?><rss xmlns:atom="http://www.w3.org/2005/Atom" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:media="http://search.yahoo.com/mrss/" version="2.0">
    <channel>
        <title>BSD Now HD</title>
        <link>http://www.jupiterbroadcasting.com/show/bsdnow/</link>
        <description>A weekly show covering the latest developments in the world of the BSD family of operating systems. News, Tutorials and Interviews for new users and long time developers alike.</description>
        <generator>Feeder 2.5.12(2294); Mac OS X Version 10.9.5 (Build 13F34) http://reinventedsoftware.com/feeder/</generator>
        <docs>http://blogs.law.harvard.edu/tech/rss</docs>
        <language>en</language>
        <pubDate>Thu, 25 Jun 2015 13:15:16 -0700</pubDate>
        <lastBuildDate>Thu, 25 Jun 2015 13:15:16 -0700</lastBuildDate>
        
        <itunes:author>Jupiter Broadcasting</itunes:author>
        <itunes:image href="http://www.jupiterbroadcasting.com/images/bsdnow-badge.jpg" />
        <itunes:explicit>no</itunes:explicit>
        
        <itunes:block>no</itunes:block>
        
        <atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="self" type="application/rss+xml" href="http://feeds.feedburner.com/BsdNowHd" /><feedburner:info xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" uri="bsdnowhd" /><atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="hub" href="http://pubsubhubbub.appspot.com/" /><media:copyright>Copyright Jupiter Broadcasting</media:copyright><media:thumbnail url="http://www.jupiterbroadcasting.com/images/bsdnow-badge.jpg" /><media:category scheme="http://www.itunes.com/dtds/podcast-1.0.dtd">Technology/Software How-To</media:category><media:category scheme="http://www.itunes.com/dtds/podcast-1.0.dtd">Technology/Tech News</media:category><itunes:owner><itunes:email>BSD, FreeBSD, PCBSD, PC-BSD, OpenBSD, NetBSD, DragonFlyBSD, FreeNAS, pfSense, Interview, Tutorial, ZFS, UFS</itunes:email><itunes:name>Jupiter Broadcasting</itunes:name></itunes:owner><itunes:subtitle>Everything you wanted to know about BSD</itunes:subtitle><itunes:summary>A weekly show covering the latest developments in the world of the BSD family of operating systems. News, Tutorials and Interviews for new users and long time developers alike.</itunes:summary><itunes:category text="Technology"><itunes:category text="Software How-To" /></itunes:category><itunes:category text="Technology"><itunes:category text="Tech News" /></itunes:category><item>
            <title>Bitrot Group Therapy | BSD Now 95</title>
            <link>http://www.jupiterbroadcasting.com/84272/bitrot-group-therapy-bsd-now-95/</link>
            <description><![CDATA[This time on the show, we'll be talking some ZFS with Sean Chittenden. He's been using it on FreeBSD at Groupon, and has some interesting stories about how it's saved his data. Answers to your emails and all of this week's headlines, on BSD Now - the place to B.. SD.]]></description>
            <pubDate>Thu, 25 Jun 2015 13:13:48 -0700</pubDate>
            <enclosure url="http://www.podtrac.com/pts/redirect.mp4/201406.jb-dl.cdn.scaleengine.net/bsdnow/2015/bsd-0095.mp4" length="510856278" type="video/mp4" />
            <guid isPermaLink="false">DD438D40-D5A1-4D08-974F-0B3FAF6BDF9D</guid>
            <itunes:author>Jupiter Broadcasting</itunes:author>
            <itunes:subtitle>This time on the show, we'll be talking some ZFS with Sean Chittenden. He's been using it on FreeBSD at Groupon, and has some interesting stories about how it's saved his data. Answers to your emails and more, on BSD Now - the place to B.. SD.</itunes:subtitle>
            <itunes:summary><![CDATA[This time on the show, we'll be talking some ZFS with Sean Chittenden. He's been using it on FreeBSD at Groupon, and has some interesting stories about how it's saved his data. Answers to your emails and all of this week's headlines, on BSD Now - the place to B.. SD.]]></itunes:summary>
            <itunes:explicit>no</itunes:explicit>
            <itunes:duration>1:15:36</itunes:duration>
            <media:thumbnail url="http://www.jupiterbroadcasting.com/wp-content/uploads/2015/06/bsd-0095-v.jpg" />
        <author>BSD, FreeBSD, PCBSD, PC-BSD, OpenBSD, NetBSD, DragonFlyBSD, FreeNAS, pfSense, Interview, Tutorial, ZFS, UFS (Jupiter Broadcasting)</author><media:content url="http://www.podtrac.com/pts/redirect.mp4/201406.jb-dl.cdn.scaleengine.net/bsdnow/2015/bsd-0095.mp4" fileSize="510856278" type="video/mp4" /></item>
        <item><guid>DD438D40-D5A1-4D08-974F-0B3FAF6BDF9C</guid></item>
        <item><title>Where dat GUID at?</title></item>
        <copyright>Copyright Jupiter Broadcasting</copyright><media:credit role="author">Jupiter Broadcasting</media:credit><media:rating>nonadult</media:rating><media:description type="plain">Everything you wanted to know about BSD</media:description></channel>
</rss>`)

var atomTestFeed = []byte(`<?xml version="1.0" encoding="utf-8"?>

<feed xmlns="http://www.w3.org/2005/Atom">

	<title>Example Feed</title>
	<subtitle>A subtitle.</subtitle>
	<link href="http://example.org/feed/" rel="self" />
	<link href="http://example.org/" />
	<id>urn:uuid:60a76c80-d399-11d9-b91C-0003939e0af6</id>
	<updated>2003-12-13T18:30:02Z</updated>
	
	
	<entry>
		<title>Atom-Powered Robots Run Amok</title>
		<link href="http://example.org/2003/12/13/atom03" />
		<link rel="alternate" type="text/html" href="http://example.org/2003/12/13/atom03.html"/>
		<link rel="edit" href="http://example.org/2003/12/13/atom03/edit"/>
		<id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</id>
		<updated>2003-12-13T18:30:02Z</updated>
		<summary>Some text.</summary>
		<content type="xhtml">
			<div xmlns="http://www.w3.org/1999/xhtml">
				<p>This is the entry content.</p>
			</div>
		</content>
		<author>
			<name>John Doe</name>
			<email>johndoe@example.com</email>
		</author>
	</entry>
	<entry>
		<title>Atom-Powered Robots Run Amok #2</title>
		<link href="http://example.org/2003/12/16/atom03" />
		<link rel="alternate" type="text/html" href="http://example.org/2003/12/13/atom03.html"/>
		<link rel="edit" href="http://example.org/2003/12/16/atom03/edit"/>
		<id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6b</id>
		<updated>2003-12-13T18:30:02Z</updated>
		<summary>Some text.</summary>
		<content type="xhtml">
			<div xmlns="http://www.w3.org/1999/xhtml">
				<p>This is the entry content.</p>
			</div>
		</content>
		<author>
			<name>John Doe</name>
			<email>johndoe@example.com</email>
		</author>
	</entry>

</feed>`)
