package console

import (
	"lego_demo/pb"
	"lego_demo/sys/db"
	"time"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/lib/modules/console"
	"github.com/liwei1dao/lego/lib/modules/http"
	"github.com/liwei1dao/lego/sys/log"
)

type ArticleApiComp struct {
	cbase.ModuleCompBase
	module console.IConsole
}

func (this *ArticleApiComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.module = module.(console.IConsole)
	api := this.module.Group("/lego/article")
	api.POST("/createarticle", this.module.CheckToken, this.CreateArticleReq)
	api.POST("/releasearticle", this.module.CheckToken, this.ReleaseArticleReq)
	api.POST("/deletearticle", this.module.CheckToken, this.DeleteArticleReq)
	api.POST("/getarticle", this.module.CheckToken, this.GetArticleReq)
	api.POST("/getauthoIrdarticles", this.module.CheckToken, this.GetAuthoIrdArticlesReq)
	api.POST("/getlastarticles", this.module.CheckToken, this.GetLastArticlesReq)
	return
}

//创建文章
func (this *ArticleApiComp) CreateArticleReq(c *http.Context) {
	uId := c.GetUInt32(console.UserKey)
	req := &pb.CreateArticleReq{}
	c.ShouldBindJSON(req)
	defer log.Debugf("CreateArticleReq:%+v", req)
	var (
		user    *console.Cache_UserData
		article *pb.DB_ArticleData
		err     error
	)
	if user, err = this.module.Cache().QueryUserData(uId); err == nil {
		article = &pb.DB_ArticleData{
			Id:           req.ArticleId,
			Title:        req.Title,
			Content:      req.Content,
			ShortContent: req.ShortContent,
			Images:       req.Images,
			CreationTime: time.Now().Unix(),
			AuthorId:     uId,
			AuthorName:   user.Db_UserData.NickName,
			AuthorAvatar: user.Db_UserData.HeadUrl,
			GreatNum:     0,
			Staet:        pb.ArticleState_Draft,
		}
		if article, err = db.CreateOrUpDateArticle(article); err == nil {
			this.module.HttpStatusOK(c, core.ErrorCode_Success, article)
		} else {
			log.Errorf("CreateArticleReq err:%v", err)
			this.module.HttpStatusOK(c, core.ErrorCode_SqlExecutionError, nil)
		}
	}
}

//发布文章
func (this *ArticleApiComp) ReleaseArticleReq(c *http.Context) {
	uId := c.GetUInt32(console.UserKey)
	req := &pb.ReleaseArticleReq{}
	c.ShouldBindJSON(req)
	defer log.Debugf("ReleaseArticleReq:%+v", req)
	var (
		user    *console.Cache_UserData
		article *pb.DB_ArticleData
		err     error
	)
	if user, err = this.module.Cache().QueryUserData(uId); err == nil {
		article = &pb.DB_ArticleData{
			Id:           req.ArticleId,
			Title:        req.Title,
			Content:      req.Content,
			ShortContent: req.ShortContent,
			Images:       req.Images,
			CreationTime: time.Now().Unix(),
			AuthorId:     uId,
			AuthorName:   user.Db_UserData.NickName,
			AuthorAvatar: user.Db_UserData.HeadUrl,
			GreatNum:     0,
			Staet:        pb.ArticleState_Review,
		}
		if article, err = db.CreateOrUpDateArticle(article); err == nil {
			this.module.HttpStatusOK(c, core.ErrorCode_Success, article)
		} else {
			log.Errorf("CreateArticleReq err:%v", err)
			this.module.HttpStatusOK(c, core.ErrorCode_SqlExecutionError, nil)
		}
	}
}

//删除文章
func (this *ArticleApiComp) DeleteArticleReq(c *http.Context) {
	uId := c.GetUInt32(console.UserKey)
	req := &pb.DeleteArticleReq{}
	c.ShouldBindJSON(req)
	defer log.Debugf("DeleteArticleReq:%+v", req)
	var (
		article *pb.DB_ArticleData
		err     error
	)
	if article, err = db.QueryArticle(req.ArticleId); err == nil && article.AuthorId == uId {
		if err = db.DeleteArticle(uId, req.ArticleId); err == nil {
			this.module.HttpStatusOK(c, core.ErrorCode_Success, article)
		} else {
			this.module.HttpStatusOK(c, core.ErrorCode_SqlExecutionError, nil)
		}
	} else {
		this.module.HttpStatusOK(c, core.ErrorCode_ReqParameterError, nil)
	}
}

//获取作家文章请求
func (this *ArticleApiComp) GetAuthoIrdArticlesReq(c *http.Context) {
	uId := c.GetUInt32(console.UserKey)
	req := &pb.GetAuthoIrdArticlesReq{}
	c.ShouldBindJSON(req)
	defer log.Debugf("GetAuthoIrdArticlesReq uId:%d - %+v ", uId, req)
	var (
		author   *pb.DB_AuthorData
		articles []*pb.DB_ArticleData
		err      error
	)
	if author, err = db.QueryAuthor(uId); err == nil {
		if articles, err = db.QueryArticles(author.ArticleIds); err == nil {
			this.module.HttpStatusOK(c, core.ErrorCode_Success, articles)
		} else {
			log.Errorf("GetArticleReq err:%v", err)
			this.module.HttpStatusOK(c, core.ErrorCode_SqlExecutionError, nil)
		}
	} else {
		log.Errorf("GetArticleReq err:%v", err)
		this.module.HttpStatusOK(c, core.ErrorCode_SqlExecutionError, nil)
	}
}

func (this *ArticleApiComp) GetArticleReq(c *http.Context) {
	req := &pb.GetArticleReq{}
	c.ShouldBindJSON(req)
	defer log.Debugf("GetArticleReq:%+v", req)
	var (
		article *pb.DB_ArticleData
		err     error
	)
	if article, err = db.QueryArticle(req.ArticleId); err == nil {
		this.module.HttpStatusOK(c, core.ErrorCode_Success, article)
	} else {
		log.Errorf("GetArticleReq err:%v", err)
		this.module.HttpStatusOK(c, core.ErrorCode_SqlExecutionError, nil)
	}
}

func (this *ArticleApiComp) GetLastArticlesReq(c *http.Context) {
	req := &pb.GetLastArticlesReq{}
	c.ShouldBindJSON(req)
	defer log.Debugf("GetLastArticlesReq:%+v", req)
	var (
		articles []*pb.DB_ArticleData
		err      error
	)
	if articles, err = db.QueryArticlesByLast(req.Start, req.Num); err == nil {
		this.module.HttpStatusOK(c, core.ErrorCode_Success, articles)
	} else {
		log.Errorf("GetLastArticlesReq err:%v", err)
		this.module.HttpStatusOK(c, core.ErrorCode_SqlExecutionError, nil)
	}
}
