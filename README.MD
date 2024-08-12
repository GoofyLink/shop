# GoFrame Template For SingleRepo

Quick Start: 
- https://goframe.org/pages/viewpage.action?pageId=1114399

# 大部分程序员在数据库中将字符集设置为 utf8mb4，排序规则设置为 utf8mb4_general_ci，主要是出于以下几个原因：
1. 全面支持 Unicode 字符：
   utf8mb4 是 MySQL 和 MariaDB 中的一个字符集，它能够支持完整的 Unicode 字符集，包括4字节的字符。标准的 utf8 只支持最多 3 字节的字符，
2. 这意味着一些特殊字符（如表情符号、某些亚洲文字、古文字等）不能被正确存储和处理。而 utf8mb4 则能够处理所有的 Unicode 字符，确保数据存储的完整性。 
   兼容性和通用性：
   utf8mb4 是一种通用的字符集，能够处理几乎所有语言和符号，这使得应用程序具有很好的国际化支持能力。即使应用程序最初不需要支持多种语言， 
   使用 utf8mb4 仍然是一个更具前瞻性的选择，因为它可以避免未来扩展时的字符集问题。
3. 排序规则选择 utf8mb4_general_ci：
   utf8mb4_general_ci 是一种不区分大小写的通用排序规则（ci 代表 case-insensitive）。这意味着在排序和比较时，不区分大小写（如 “a” 和 “A” 被视为相等），
   同时也忽略一些特殊字符的差异。这种规则在大多数情况下都是足够的，并且执行效率比更严格的排序规则要高。
   选择 utf8mb4_general_ci 可以满足大多数应用场景下的排序需求，并且比一些更加严格的排序规则（如 utf8mb4_unicode_ci）在性能上有所优势，因为它的实现更简单。
4. 性能与存储的平衡：
   尽管 utf8mb4 在存储上可能会比 utf8 多消耗一些空间（因为每个字符最多使用 4 个字节，而 utf8 最多使用 3 个字节），
   但这是为了支持更广泛的字符集所做的权衡。大多数现代应用程序都会优先考虑兼容性和功能的完整性，因此选择 utf8mb4 是一种普遍的做法。
5. 最佳实践和社区推荐：
   utf8mb4 和 utf8mb4_general_ci 已经成为了行业的最佳实践和标准配置，广泛被社区和文档推荐。这种广泛使用也确保了其稳定性和可靠性。
   总结
   使用 utf8mb4 字符集和 utf8mb4_general_ci 排序规则是一种确保数据库能够处理所有 Unicode 字符，具有国际化支持，且在大多数情况下具备良好性能的最佳实践。
   因此，大部分程序员会选择这种配置。

# 商品中心相关表结构
goods_info 商品  有goods_id
goods_options_info 商品规格 与goods_id关联
# 间接相关
## 商品中心之前开发
category_info 商品分类 parent_id 父级Id
consignee_info 收货地址
coupon_info 优惠券
## 商品中心之前开发
collection_info 收藏 user_id object_id
comment_info 评价
praise_info  点赞
