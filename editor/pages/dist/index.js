"use strict";
exports.__esModule = true;
var yaml_editor_1 = require("components/yaml-editor");
var react_1 = require("react");
var react_redux_1 = require("react-redux");
var rsuite_1 = require("rsuite");
var loquiapi_1 = require("services/loquiapi");
var bot_1 = require("store/bot");
var index_1 = require("../components/side-panel/index");
var Home = function () {
    var botId = react_redux_1.useSelector(bot_1.selectActiveBotId);
    var _a = loquiapi_1.useGetBotYamlQuery(botId), data = _a.data, isFetching = _a.isFetching, isLoading = _a.isLoading;
    if (isFetching || isLoading)
        return react_1["default"].createElement(rsuite_1.Loader, { backdrop: true, inverse: true, center: true, content: 'loading...', vertical: true });
    return (react_1["default"].createElement("div", { className: 'grid place-items-center h-screen' },
        react_1["default"].createElement(rsuite_1.Container, { className: 'place-items-center w-full h-full', style: { height: '100vh' } },
            react_1["default"].createElement(rsuite_1.Sidebar, null,
                react_1["default"].createElement(index_1["default"], null)),
            react_1["default"].createElement(rsuite_1.Container, null,
                react_1["default"].createElement(yaml_editor_1["default"], { data: data, botId: botId })))));
};
exports["default"] = Home;
