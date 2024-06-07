const express=require('express');
const router=express.Router();
const rssParser=require('rss-parser');
const Feed=require('../models/Feed');

router.get('/fetch-feed',async (req,res)=>{
    const parser=new rssParser;
    const feedURL=req.query.url;
    try {
        const feed= await parser.parseURL(feedURL);
        const newFeed=new Feed({
            title:feed.title,
            url:feedURL,
        })
        await newFeed.save();
        res.json(feed);
    } catch (err) {
        console.error(err)
        res.status(500).json({error:'failed to parse rss feed'})
    }
})