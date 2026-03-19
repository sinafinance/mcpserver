# 产品名称

新浪财经

# 产品定位

新浪财经 MCP 服务致力于成为金融领域大模型与外部数据源、工具之间的高效桥梁，通过标准化的协议，让金融数据的获取与分析更加流畅，为金融机构、投资者、金融科技开发者等提供一站式的智能金融数据交互及策略构建支持，助力金融行业智能化升级。

# 核心功能

资讯数据：新闻快讯、公司公告、研报摘要、事件脉络
金融基础：A股、港股、美股等市场静态信息、行业｜概念｜指数成分
市场行情：多品种市场行情
财务数据：财务指标
特色数据：舆情分析、持仓估算、融资融券、北向资金
衍生支持：大事提醒（各种大事数据）

# 适用人群

1) 资深投资人，可借助MCP工具高效解决自己的个性化投资问题。
2) 投资顾问，可借助工具搭建智能体，自用或为客户提供专业智能的金融服务。
3) 更多金融从业人员包含开发者、业务人员，可借助工具快速组装、搭建更多高质量的金融服务。

# 常见问题解答

1. **工具是否收费？**
   详见新浪财经芝麻开门智研平台https://zyhub.finance.sina.cn/help/index
2. **接入方式**
   详见新浪财经芝麻开门智研平台：https://zyhub.finance.sina.cn/help/interface
3. **使用 新浪财经 MCP 能否进行商用？**
   新浪财经的MCP和API不支持商用。
4. **哪里可以用到 新浪财经 MCP 工具？**
   所有支持MCP协议的AI应用、Agent平台都可以接入使用。比如Cursor、Cline、CherryStudio、Windsurf等。


# 工具列表

展示我的服务页面的工具表格 https://zyhub.finance.sina.cn/myservice/index

| 工具  | 描述  |
| --- | --- |
| cnSectorComponentsRanking | A股指数和行业板块成分股 |
| globalStockQuoteRealtime | 股票行情 |
| globalStockKlineDaily | 全市场股票日K |
| globalStockHotBoard | 股票热搜榜 |
| cnStockRatingHistory | A股上市公司信披考评 |
| usSectorRanking | 美股板块列表 |
| usMarketStatisticsUpdown | 美股市场涨跌分布 |
| globalStockMajorEvents | 全市场公司重大事项 |
| hkStockHotBoard | 港股热度榜 |
| hkSectorLeadingStocks | 领涨板块的个股列表 |
| forexQuotesBatch | 批量外汇行情 |
| cnMarketUpdownDistribution | A股全市场涨跌分布统计 |
| cnCompanyBasicInfo | A股公司基础信息 |
| cnStockKLine | A股股票分钟K线 |
| cnFinanceReportsFull | A股公司财务指标 |
| cnStockValuationDetail | A股公司估值数据 |
| usCompanyManagerInfo | 美国市场公司高管 |
| usTradingFundFlow1Day | 美股股今日资金趋势 |
| usTradingFundFlow5Days | 美股5日累计净流入 |
| usTradingFundFlow60Days | 美股股主力资金历史 |
| hkSectorQuotesList | 港股板块行情列表 |
| cnFinanceRevenueComposition | A股公司主营构成 |
| futureCommodityList | 国内商品期货 |
| futureFinancialList | 金融期货 |
| futureGoldList | 黄金综合期货 |
| cnCompanyManagerInfo | A股公司高管详情 |
| cnCompanyCapitalHistory | A股公司股本变动 |
| cnCompanyShareholderHistory | A股公司股东人数 |
| hkTradingMainFundsHistory | 港股主力资金历史 |
| cnVirtualSectorRanking | A股虚拟板块列表 |
| hkFinanceReportsByIndex | 港股财务指标 |
| cnTradingBlockList | A股大宗交易 |
| cnStockLockupFuture | A股股票未来解禁 |
| cnStockLockupHistory | A股股票历史解禁 |
| cnStockTradingMarginList | A股股票融资融券 |
| cnMarketLimitUpPool | A股市场涨停池 |
| cnStockLianBC | A股聚焦连板个股 |
| cnMarketStrongSectors | A股强势板块列表 |
| hkStockSpecialRanking | 港股特色榜单 |
| cnStockConnectHoldings | A股沪深港通 |
| globalStockSearchSymbols | 全市场股票代码搜索接口 |
| newsArticleDetail | 单篇资讯详情 |
| newsArticleList | 栏目资讯列表 |
| newsFlashDetail | 快讯详情 |
| newsFlashList | 快讯列表 |
| forexQuoteLatest | 获取最新汇率行情 |
| cnFinanceReportDateList | A股公司财报日期列表 |
| cnFinanceGrReportDateList | A股公司主营构成报告日期列表 |

## MCP Server 配置

- **认证方式**
  
    MCP Server连接认证支持以下两种认证方式，注意两种方式使用时候的**字段名称区别**，使用时候将 **\<TOKEN\>** 替换为申请的token
  
  - query参数认证：

```HTML
https://mcp.finance.sina.com.cn/sse?token=<TOKEN>
```

- header携带认证

```HTML
{
    "headers": {
        "X-Auth-Token": "<TOKEN>"
    }
}
```

- 协议支持

MCP Server支持以下两种连接协议配置，注意以下例子以 **cursor 2.1.39** 版本为验证对象

- SSE 协议使用配置

```JSON
# header 携带认证参数
{
    "mcpServers": {
        "sinafinance-sse-test": {
            "type": "sse",
            "url": "http://mcp.finance.sina.com.cn/sse",
            "headers": {
                "X-Auth-Token": "<TOKEN>"
            }
        }
    }
}

# query 参数携带认证信息
{
    "mcpServers": {
        "sinafinance-sse-test": {
            "type": "sse",
            "url": "http://mcp.finance.sina.com.cn/sse?token=<TOKEN>",
        }
    }
}
```

- Streamable Http协议(推荐)使用配置

```JSON
# header 认证方式
{
    "mcpServers": {
        "sinafinance-http-test": {
            "type": "streamable-http",
            "url": "http://mcp.finance.sina.com.cn/mcp-http",
            "name": "sinafinance-http",
            "headers": {
                "X-Auth-Token": "<TOKEN>"
            }
        }
    }
}


# query 认证
{
    "mcpServers": {
        "sinafinance-http-test": {
            "type": "streamable-http",
            "url": "http://mcp.finance.sina.com.cn/mcp-http?token=<TOKEN>",
            "name": "sinafinance-http"
        }
    }
}
```

### 设置说明

您可以使用Cursor、Dify、Cline、CherryStudio或其他客户端来连接到我们的MCP Server。
