module.exports = {
    plugins: [
        "@xyclos/references"
    ],
    rules: {
        "type-enum": [2, "always", [
            "fix", // 表示在代码库中修复了一个 bug（这和语义化版本中的 PATCH 相对应）
            "feat", // 表示在代码库中新增了一个功能（这和语义化版本中的 MINOR 相对应）
            "docs", // 用于修改文档，例如修改 README 文件、API 文档等
            "style", // 用于修改代码的样式，例如调整缩进、空格、空行等
            "refactor", // 用于重构代码，例如修改代码结构、变量名、函数名等但不修改功能逻辑
            "perf", // 用于优化性能，例如提升代码的性能、减少内存占用等
            "test", // 用于修改测试用例，例如添加、删除、修改代码的测试用例等
            "revert", // 用于撤销之前的 commit
            "chore", // 用户修改构建过程或辅助工具，例如修改 Makefile 文件、增加依赖库等
            "build", // 用于修改项目构建系统，例如修改依赖库、外部接口或者升级 Node 版本等
            "ci" // 用于修改持续集成流程，例如修改 Travis、Jenkins 等工作流配置
        ]],
        "references-empty-enum": [2, "never", [
            "fix",
            "feat"
        ]]
    },
    parserPreset: {
        parserOpts: {
            issuePrefixes: ["#"]
        }
    },
    helpUrl: 'https://github.com/conventional-changelog/commitlint/#what-is-commitlint',
}