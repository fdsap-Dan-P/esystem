eval("/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__,
\"a\", function() { return MODEL_PROP_NAME; });\n/* unused harmony export
MODEL_EVENT_NAME */\n/* harmony export (binding) */
__webpack_require__.d(__webpack_exports__, \"c\", function() { return props;
});\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__,
\"b\", function() { return paginationMixin; });\n/* harmony import */ var
__WEBPACK_IMPORTED_MODULE_0__vue__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/vue.js\");\n/* harmony
import */ var __WEBPACK_IMPORTED_MODULE_1__constants_components__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/constants/components.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_2__constants_key_codes__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/constants/key-codes.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_3__constants_props__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/constants/props.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_4__constants_slots__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/constants/slots.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_5__utils_array__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/array.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_6__utils_dom__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/dom.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_7__utils_events__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/events.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_8__utils_inspect__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/inspect.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_9__utils_math__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/math.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_10__utils_model__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/model.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_11__utils_number__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/number.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_12__utils_object__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/object.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_13__utils_props__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/props.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_14__utils_safe_vue_instance__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/safe-vue-instance.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_15__utils_string__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/string.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_16__utils_warn__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/utils/warn.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_17__mixins_normalize_slot__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/mixins/normalize-slot.js\");\n/*
harmony import */ var __WEBPACK_IMPORTED_MODULE_18__components_link_link__ =
__webpack_require__(\"./node_modules/bootstrap-vue/esm/components/link/link.js\");\nvar
_watch;\nfunction ownKeys(object, enumerableOnly) {\n var keys =
Object.keys(object);\n if (Object.getOwnPropertySymbols) {\n var symbols =
Object.getOwnPropertySymbols(object);\n enumerableOnly && (symbols =
symbols.filter(function (sym) {\n return Object.getOwnPropertyDescriptor(object,
sym).enumerable;\n })), keys.push.apply(keys, symbols);\n }\n return
keys;\n}\nfunction _objectSpread(target) {\n for (var i = 1; i <
arguments.length; i++) {\n var source = null != arguments[i] ? arguments[i] :
{};\n i % 2 ? ownKeys(Object(source), !0).forEach(function (key) {\n
_defineProperty(target, key, source[key]);\n }) :
Object.getOwnPropertyDescriptors ? Object.defineProperties(target,
Object.getOwnPropertyDescriptors(source)) :
ownKeys(Object(source)).forEach(function (key) {\n Object.defineProperty(target,
key, Object.getOwnPropertyDescriptor(source, key));\n });\n }\n return
target;\n}\nfunction _defineProperty(obj, key, value) {\n if (key in obj) {\n
Object.defineProperty(obj, key, {\n value: value,\n enumerable: true,\n
configurable: true,\n writable: true\n });\n } else {\n obj[key] = value;\n }\n
return obj;\n}\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n // Common props, computed,
data, render function, and methods\n// for `<b-pagination
  >` and `<b-pagination-nav
    >`\n// --- Constants ---\n\nvar _makeModelMixin =
    Object(__WEBPACK_IMPORTED_MODULE_10__utils_model__[\"a\" /* makeModelMixin
    */])('value', {\n type: __WEBPACK_IMPORTED_MODULE_3__constants_props__[\"i\"
    /* PROP_TYPE_BOOLEAN_NUMBER_STRING */],\n defaultValue: null,\n /* istanbul
    ignore next */\n validator: function validator(value) {\n if
    (!Object(__WEBPACK_IMPORTED_MODULE_8__utils_inspect__[\"g\" /* isNull
    */])(value) && Object(__WEBPACK_IMPORTED_MODULE_11__utils_number__[\"c\" /*
    toInteger */])(value, 0) < 1) {\n
    Object(__WEBPACK_IMPORTED_MODULE_16__utils_warn__[\"a\" /* warn
    */])('\"v-model\" value must be a number greater than \"0\"',
    __WEBPACK_IMPORTED_MODULE_1__constants_components__[\"_40\" /*
    NAME_PAGINATION */]);\n return false;\n }\n return true;\n }\n }),\n
    modelMixin = _makeModelMixin.mixin,\n modelProps = _makeModelMixin.props,\n
    MODEL_PROP_NAME = _makeModelMixin.prop,\n MODEL_EVENT_NAME =
    _makeModelMixin.event;\n // Threshold of limit size when we start/stop
    showing ellipsis\n\nvar ELLIPSIS_THRESHOLD = 3; // Default # of buttons
    limit\n\nvar DEFAULT_LIMIT = 5; // --- Helper methods ---\n// Make an array
    of N to N+X\n\nvar makePageArray = function makePageArray(startNumber,
    numberOfPages) {\n return
    Object(__WEBPACK_IMPORTED_MODULE_5__utils_array__[\"c\" /* createArray
    */])(numberOfPages, function (_, i) {\n return {\n number: startNumber +
    i,\n classes: null\n };\n });\n}; // Sanitize the provided limit value
    (converting to a number)\n\nvar sanitizeLimit = function
    sanitizeLimit(value) {\n var limit =
    Object(__WEBPACK_IMPORTED_MODULE_11__utils_number__[\"c\" /* toInteger
    */])(value) || 1;\n return limit < 1 ? DEFAULT_LIMIT : limit;\n}; //
    Sanitize the provided current page number (converting to a number)\n\nvar
    sanitizeCurrentPage = function sanitizeCurrentPage(val, numberOfPages) {\n
    var page = Object(__WEBPACK_IMPORTED_MODULE_11__utils_number__[\"c\" /*
    toInteger */])(val) || 1;\n return page > numberOfPages ? numberOfPages :
    page < 1 ? 1 : page;\n}; // Links don't normally respond to SPACE, so we add
    that\n// functionality via this handler\n\nvar onSpaceKey = function
    onSpaceKey(event) {\n if (event.keyCode ===
    __WEBPACK_IMPORTED_MODULE_2__constants_key_codes__[\"l\" /* CODE_SPACE */])
    {\n // Stop page from scrolling\n
    Object(__WEBPACK_IMPORTED_MODULE_7__utils_events__[\"f\" /* stopEvent
    */])(event, {\n immediatePropagation: true\n }); // Trigger the click event
    on the link\n\n event.currentTarget.click();\n return false;\n }\n}; // ---
    Props ---\n\nvar props =
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"d\" /*
    makePropsConfigurable
    */])(Object(__WEBPACK_IMPORTED_MODULE_12__utils_object__[\"m\" /* sortKeys
    */])(_objectSpread(_objectSpread({}, modelProps), {}, {\n align:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], 'left'),\n ariaLabel:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], 'Pagination'),\n disabled:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"g\" /*
    PROP_TYPE_BOOLEAN */], false),\n ellipsisClass:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"e\" /*
    PROP_TYPE_ARRAY_OBJECT_STRING */]),\n ellipsisText:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], \"\\u2026\"),\n // '…'\n firstClass:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"e\" /*
    PROP_TYPE_ARRAY_OBJECT_STRING */]),\n firstNumber:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"g\" /*
    PROP_TYPE_BOOLEAN */], false),\n firstText:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], \"\\xAB\"),\n // '«'\n hideEllipsis:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"g\" /*
    PROP_TYPE_BOOLEAN */], false),\n hideGotoEndButtons:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"g\" /*
    PROP_TYPE_BOOLEAN */], false),\n labelFirstPage:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], 'Go to first page'),\n labelLastPage:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], 'Go to last page'),\n labelNextPage:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], 'Go to next page'),\n labelPage:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"m\" /*
    PROP_TYPE_FUNCTION_STRING */], 'Go to page'),\n labelPrevPage:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], 'Go to previous page'),\n lastClass:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"e\" /*
    PROP_TYPE_ARRAY_OBJECT_STRING */]),\n lastNumber:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"g\" /*
    PROP_TYPE_BOOLEAN */], false),\n lastText:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], \"\\xBB\"),\n // '»'\n limit:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"p\" /*
    PROP_TYPE_NUMBER_STRING */], DEFAULT_LIMIT, /* istanbul ignore next */\n
    function (value) {\n if
    (Object(__WEBPACK_IMPORTED_MODULE_11__utils_number__[\"c\" /* toInteger
    */])(value, 0) < 1) {\n
    Object(__WEBPACK_IMPORTED_MODULE_16__utils_warn__[\"a\" /* warn */])('Prop
    \"limit\" must be a number greater than \"0\"',
    __WEBPACK_IMPORTED_MODULE_1__constants_components__[\"_40\" /*
    NAME_PAGINATION */]);\n return false;\n }\n return true;\n }),\n nextClass:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"e\" /*
    PROP_TYPE_ARRAY_OBJECT_STRING */]),\n nextText:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], \"\\u203A\"),\n // '›'\n pageClass:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"e\" /*
    PROP_TYPE_ARRAY_OBJECT_STRING */]),\n pills:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"g\" /*
    PROP_TYPE_BOOLEAN */], false),\n prevClass:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"e\" /*
    PROP_TYPE_ARRAY_OBJECT_STRING */]),\n prevText:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */], \"\\u2039\"),\n // '‹'\n size:
    Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"c\" /* makeProp
    */])(__WEBPACK_IMPORTED_MODULE_3__constants_props__[\"u\" /*
    PROP_TYPE_STRING */])\n})), 'pagination'); // --- Mixin ---\n//
    @vue/component\n\nvar paginationMixin =
    Object(__WEBPACK_IMPORTED_MODULE_0__vue__[\"d\" /* extend */])({\n mixins:
    [modelMixin, __WEBPACK_IMPORTED_MODULE_17__mixins_normalize_slot__[\"a\" /*
    normalizeSlotMixin */]],\n props: props,\n data: function data() {\n // `-1`
    signifies no page initially selected\n var currentPage =
    Object(__WEBPACK_IMPORTED_MODULE_11__utils_number__[\"c\" /* toInteger
    */])(this[MODEL_PROP_NAME], 0);\n currentPage = currentPage > 0 ?
    currentPage : -1;\n return {\n currentPage: currentPage,\n
    localNumberOfPages: 1,\n localLimit: DEFAULT_LIMIT\n };\n },\n computed: {\n
    btnSize: function btnSize() {\n var size = this.size;\n return size ?
    \"pagination-\".concat(size) : '';\n },\n alignment: function alignment()
    {\n var align = this.align;\n if (align === 'center') {\n return
    'justify-content-center';\n } else if (align === 'end' || align === 'right')
    {\n return 'justify-content-end';\n } else if (align === 'fill') {\n // The
    page-items will also have 'flex-fill' added\n // We add text centering to
    make the button appearance better in fill mode\n return 'text-center';\n }\n
    return '';\n },\n styleClass: function styleClass() {\n return this.pills ?
    'b-pagination-pills' : '';\n },\n computedCurrentPage: function
    computedCurrentPage() {\n return sanitizeCurrentPage(this.currentPage,
    this.localNumberOfPages);\n },\n paginationParams: function
    paginationParams() {\n // Determine if we should show the the ellipsis\n var
    limit = this.localLimit,\n numberOfPages = this.localNumberOfPages,\n
    currentPage = this.computedCurrentPage,\n hideEllipsis =
    this.hideEllipsis,\n firstNumber = this.firstNumber,\n lastNumber =
    this.lastNumber;\n var showFirstDots = false;\n var showLastDots = false;\n
    var numberOfLinks = limit;\n var startNumber = 1;\n if (numberOfPages <=
    limit) {\n // Special case: Less pages available than the limit of displayed
    pages\n numberOfLinks = numberOfPages;\n } else if (currentPage < limit - 1
    && limit > ELLIPSIS_THRESHOLD) {\n if (!hideEllipsis || lastNumber) {\n
    showLastDots = true;\n numberOfLinks = limit - (firstNumber ? 0 : 1);\n }\n
    numberOfLinks = Object(__WEBPACK_IMPORTED_MODULE_9__utils_math__[\"e\" /*
    mathMin */])(numberOfLinks, limit);\n } else if (numberOfPages - currentPage
    + 2 < limit && limit > ELLIPSIS_THRESHOLD) {\n if (!hideEllipsis ||
    firstNumber) {\n showFirstDots = true;\n numberOfLinks = limit - (lastNumber
    ? 0 : 1);\n }\n startNumber = numberOfPages - numberOfLinks + 1;\n } else
    {\n // We are somewhere in the middle of the page list\n if (limit >
    ELLIPSIS_THRESHOLD) {\n numberOfLinks = limit - (hideEllipsis ? 0 : 2);\n
    showFirstDots = !!(!hideEllipsis || firstNumber);\n showLastDots =
    !!(!hideEllipsis || lastNumber);\n }\n startNumber = currentPage -
    Object(__WEBPACK_IMPORTED_MODULE_9__utils_math__[\"c\" /* mathFloor
    */])(numberOfLinks / 2);\n } // Sanity checks\n\n /* istanbul ignore if
    */\n\n if (startNumber < 1) {\n startNumber = 1;\n showFirstDots = false;\n
    } else if (startNumber > numberOfPages - numberOfLinks) {\n startNumber =
    numberOfPages - numberOfLinks + 1;\n showLastDots = false;\n }\n if
    (showFirstDots && firstNumber && startNumber < 4) {\n numberOfLinks =
    numberOfLinks + 2;\n startNumber = 1;\n showFirstDots = false;\n }\n var
    lastPageNumber = startNumber + numberOfLinks - 1;\n if (showLastDots &&
    lastNumber && lastPageNumber > numberOfPages - 3) {\n numberOfLinks =
    numberOfLinks + (lastPageNumber === numberOfPages - 2 ? 2 : 3);\n
    showLastDots = false;\n } // Special handling for lower limits (where
    ellipsis are never shown)\n\n if (limit <= ELLIPSIS_THRESHOLD) {\n if
    (firstNumber && startNumber === 1) {\n numberOfLinks =
    Object(__WEBPACK_IMPORTED_MODULE_9__utils_math__[\"e\" /* mathMin
    */])(numberOfLinks + 1, numberOfPages, limit + 1);\n } else if (lastNumber
    && numberOfPages === startNumber + numberOfLinks - 1) {\n startNumber =
    Object(__WEBPACK_IMPORTED_MODULE_9__utils_math__[\"d\" /* mathMax
    */])(startNumber - 1, 1);\n numberOfLinks =
    Object(__WEBPACK_IMPORTED_MODULE_9__utils_math__[\"e\" /* mathMin
    */])(numberOfPages - startNumber + 1, numberOfPages, limit + 1);\n }\n }\n
    numberOfLinks = Object(__WEBPACK_IMPORTED_MODULE_9__utils_math__[\"e\" /*
    mathMin */])(numberOfLinks, numberOfPages - startNumber + 1);\n return {\n
    showFirstDots: showFirstDots,\n showLastDots: showLastDots,\n numberOfLinks:
    numberOfLinks,\n startNumber: startNumber\n };\n },\n pageList: function
    pageList() {\n // Generates the pageList array\n var _this$paginationParam =
    this.paginationParams,\n numberOfLinks =
    _this$paginationParam.numberOfLinks,\n startNumber =
    _this$paginationParam.startNumber;\n var currentPage =
    this.computedCurrentPage; // Generate list of page numbers\n\n var pages =
    makePageArray(startNumber, numberOfLinks); // We limit to a total of 3 page
    buttons on XS screens\n // So add classes to page links to hide them for XS
    breakpoint\n // Note: Ellipsis will also be hidden on XS screens\n // TODO:
    Make this visual limit configurable based on breakpoint(s)\n\n if
    (pages.length > 3) {\n var idx = currentPage - startNumber; // THe following
    is a bootstrap-vue custom utility class\n\n var classes =
    'bv-d-xs-down-none';\n if (idx === 0) {\n // Keep leftmost 3 buttons visible
    when current page is first page\n for (var i = 3; i < pages.length; i++) {\n
    pages[i].classes = classes;\n }\n } else if (idx === pages.length - 1) {\n
    // Keep rightmost 3 buttons visible when current page is last page\n for
    (var _i = 0; _i < pages.length - 3; _i++) {\n pages[_i].classes = classes;\n
    }\n } else {\n // Hide all except current page, current page - 1 and current
    page + 1\n for (var _i2 = 0; _i2 < idx - 1; _i2++) {\n // hide some left
    button(s)\n pages[_i2].classes = classes;\n }\n for (var _i3 = pages.length
    - 1; _i3 > idx + 1; _i3--) {\n // hide some right button(s)\n
    pages[_i3].classes = classes;\n }\n }\n }\n return pages;\n }\n },\n watch:
    (_watch = {}, _defineProperty(_watch, MODEL_PROP_NAME, function (newValue,
    oldValue) {\n if (newValue !== oldValue) {\n this.currentPage =
    sanitizeCurrentPage(newValue, this.localNumberOfPages);\n }\n }),
    _defineProperty(_watch, \"currentPage\", function currentPage(newValue,
    oldValue) {\n if (newValue !== oldValue) {\n // Emit `null` if no page
    selected\n this.$emit(MODEL_EVENT_NAME, newValue > 0 ? newValue : null);\n
    }\n }), _defineProperty(_watch, \"limit\", function limit(newValue,
    oldValue) {\n if (newValue !== oldValue) {\n this.localLimit =
    sanitizeLimit(newValue);\n }\n }), _watch),\n created: function created()
    {\n var _this = this;\n\n // Set our default values in data\n
    this.localLimit = sanitizeLimit(this.limit);\n this.$nextTick(function ()
    {\n // Sanity check\n _this.currentPage = _this.currentPage >
    _this.localNumberOfPages ? _this.localNumberOfPages : _this.currentPage;\n
    });\n },\n methods: {\n handleKeyNav: function handleKeyNav(event) {\n var
    keyCode = event.keyCode,\n shiftKey = event.shiftKey;\n /* istanbul ignore
    if */\n\n if (this.isNav) {\n // We disable left/right keyboard navigation
    in `<b-pagination-nav
      >`\n return;\n }\n if (keyCode ===
      __WEBPACK_IMPORTED_MODULE_2__constants_key_codes__[\"h\" /* CODE_LEFT */]
      || keyCode === __WEBPACK_IMPORTED_MODULE_2__constants_key_codes__[\"m\" /*
      CODE_UP */]) {\n Object(__WEBPACK_IMPORTED_MODULE_7__utils_events__[\"f\"
      /* stopEvent */])(event, {\n propagation: false\n });\n shiftKey ?
      this.focusFirst() : this.focusPrev();\n } else if (keyCode ===
      __WEBPACK_IMPORTED_MODULE_2__constants_key_codes__[\"k\" /* CODE_RIGHT */]
      || keyCode === __WEBPACK_IMPORTED_MODULE_2__constants_key_codes__[\"c\" /*
      CODE_DOWN */]) {\n
      Object(__WEBPACK_IMPORTED_MODULE_7__utils_events__[\"f\" /* stopEvent
      */])(event, {\n propagation: false\n });\n shiftKey ? this.focusLast() :
      this.focusNext();\n }\n },\n getButtons: function getButtons() {\n //
      Return only buttons that are visible\n return
      Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"F\" /* selectAll
      */])('button.page-link, a.page-link', this.$el).filter(function (btn) {\n
      return Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"u\" /* isVisible
      */])(btn);\n });\n },\n focusCurrent: function focusCurrent() {\n var
      _this2 = this;\n\n // We do this in `$nextTick()` to ensure buttons have
      finished rendering\n this.$nextTick(function () {\n var btn =
      _this2.getButtons().find(function (el) {\n return
      Object(__WEBPACK_IMPORTED_MODULE_11__utils_number__[\"c\" /* toInteger
      */])(Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"h\" /* getAttr
      */])(el, 'aria-posinset'), 0) === _this2.computedCurrentPage;\n });\n if
      (!Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"d\" /* attemptFocus
      */])(btn)) {\n // Fallback if current page is not in button list\n
      _this2.focusFirst();\n }\n });\n },\n focusFirst: function focusFirst()
      {\n var _this3 = this;\n\n // We do this in `$nextTick()` to ensure
      buttons have finished rendering\n this.$nextTick(function () {\n var btn =
      _this3.getButtons().find(function (el) {\n return
      !Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"r\" /* isDisabled
      */])(el);\n });\n Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"d\" /*
      attemptFocus */])(btn);\n });\n },\n focusLast: function focusLast() {\n
      var _this4 = this;\n\n // We do this in `$nextTick()` to ensure buttons
      have finished rendering\n this.$nextTick(function () {\n var btn =
      _this4.getButtons().reverse().find(function (el) {\n return
      !Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"r\" /* isDisabled
      */])(el);\n });\n Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"d\" /*
      attemptFocus */])(btn);\n });\n },\n focusPrev: function focusPrev() {\n
      var _this5 = this;\n\n // We do this in `$nextTick()` to ensure buttons
      have finished rendering\n this.$nextTick(function () {\n var buttons =
      _this5.getButtons();\n var index =
      buttons.indexOf(Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"g\" /*
      getActiveElement */])());\n if (index > 0 &&
      !Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"r\" /* isDisabled
      */])(buttons[index - 1])) {\n
      Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"d\" /* attemptFocus
      */])(buttons[index - 1]);\n }\n });\n },\n focusNext: function focusNext()
      {\n var _this6 = this;\n\n // We do this in `$nextTick()` to ensure
      buttons have finished rendering\n this.$nextTick(function () {\n var
      buttons = _this6.getButtons();\n var index =
      buttons.indexOf(Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"g\" /*
      getActiveElement */])());\n if (index < buttons.length - 1 &&
      !Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"r\" /* isDisabled
      */])(buttons[index + 1])) {\n
      Object(__WEBPACK_IMPORTED_MODULE_6__utils_dom__[\"d\" /* attemptFocus
      */])(buttons[index + 1]);\n }\n });\n }\n },\n render: function render(h)
      {\n var _this7 = this;\n var _safeVueInstance =
      Object(__WEBPACK_IMPORTED_MODULE_14__utils_safe_vue_instance__[\"a\" /*
      safeVueInstance */])(this),\n disabled = _safeVueInstance.disabled,\n
      labelPage = _safeVueInstance.labelPage,\n ariaLabel =
      _safeVueInstance.ariaLabel,\n isNav = _safeVueInstance.isNav,\n
      numberOfPages = _safeVueInstance.localNumberOfPages,\n currentPage =
      _safeVueInstance.computedCurrentPage;\n var pageNumbers =
      this.pageList.map(function (p) {\n return p.number;\n });\n var
      _this$paginationParam2 = this.paginationParams,\n showFirstDots =
      _this$paginationParam2.showFirstDots,\n showLastDots =
      _this$paginationParam2.showLastDots;\n var fill = this.align === 'fill';\n
      var $buttons = []; // Helper function and flag\n\n var isActivePage =
      function isActivePage(pageNumber) {\n return pageNumber === currentPage;\n
      };\n var noCurrentPage = this.currentPage < 1; // Factory function for
      prev/next/first/last buttons\n\n var makeEndBtn = function
      makeEndBtn(linkTo, ariaLabel, btnSlot, btnText, btnClass, pageTest, key)
      {\n var isDisabled = disabled || isActivePage(pageTest) || noCurrentPage
      || linkTo < 1 || linkTo > numberOfPages;\n var pageNumber = linkTo < 1 ? 1
      : linkTo > numberOfPages ? numberOfPages : linkTo;\n var scope = {\n
      disabled: isDisabled,\n page: pageNumber,\n index: pageNumber - 1\n };\n
      var $btnContent = _this7.normalizeSlot(btnSlot, scope) ||
      Object(__WEBPACK_IMPORTED_MODULE_15__utils_string__[\"g\" /* toString
      */])(btnText) || h();\n var $inner = h(isDisabled ? 'span' : isNav ?
      __WEBPACK_IMPORTED_MODULE_18__components_link_link__[\"a\" /* BLink */] :
      'button', {\n staticClass: 'page-link',\n class: {\n 'flex-grow-1': !isNav
      && !isDisabled && fill\n },\n props: isDisabled || !isNav ? {} :
      _this7.linkProps(linkTo),\n attrs: {\n role: isNav ? null : 'menuitem',\n
      type: isNav || isDisabled ? null : 'button',\n tabindex: isDisabled ||
      isNav ? null : '-1',\n 'aria-label': ariaLabel,\n 'aria-controls':
      Object(__WEBPACK_IMPORTED_MODULE_14__utils_safe_vue_instance__[\"a\" /*
      safeVueInstance */])(_this7).ariaControls || null,\n 'aria-disabled':
      isDisabled ? 'true' : null\n },\n on: isDisabled ? {} : {\n '!click':
      function click(event) {\n _this7.onClick(event, linkTo);\n },\n keydown:
      onSpaceKey\n }\n }, [$btnContent]);\n return h('li', {\n key: key,\n
      staticClass: 'page-item',\n class: [{\n disabled: isDisabled,\n
      'flex-fill': fill,\n 'd-flex': fill && !isNav && !isDisabled\n },
      btnClass],\n attrs: {\n role: isNav ? null : 'presentation',\n
      'aria-hidden': isDisabled ? 'true' : null\n }\n }, [$inner]);\n }; //
      Ellipsis factory\n\n var makeEllipsis = function makeEllipsis(isLast) {\n
      return h('li', {\n staticClass: 'page-item',\n class: ['disabled',
      'bv-d-xs-down-none', fill ? 'flex-fill' : '', _this7.ellipsisClass],\n
      attrs: {\n role: 'separator'\n },\n key: \"ellipsis-\".concat(isLast ?
      'last' : 'first')\n }, [h('span', {\n staticClass: 'page-link'\n },
      [_this7.normalizeSlot(__WEBPACK_IMPORTED_MODULE_4__constants_slots__[\"m\"
      /* SLOT_NAME_ELLIPSIS_TEXT */]) ||
      Object(__WEBPACK_IMPORTED_MODULE_15__utils_string__[\"g\" /* toString
      */])(_this7.ellipsisText) || h()])]);\n }; // Page button factory\n\n var
      makePageButton = function makePageButton(page, idx) {\n var pageNumber =
      page.number;\n var active = isActivePage(pageNumber) && !noCurrentPage; //
      Active page will have tabindex of 0, or if no current page and first page
      button\n\n var tabIndex = disabled ? null : active || noCurrentPage && idx
      === 0 ? '0' : '-1';\n var attrs = {\n role: isNav ? null :
      'menuitemradio',\n type: isNav || disabled ? null : 'button',\n
      'aria-disabled': disabled ? 'true' : null,\n 'aria-controls':
      Object(__WEBPACK_IMPORTED_MODULE_14__utils_safe_vue_instance__[\"a\" /*
      safeVueInstance */])(_this7).ariaControls || null,\n 'aria-label':
      Object(__WEBPACK_IMPORTED_MODULE_13__utils_props__[\"b\" /*
      hasPropFunction */])(labelPage) ? /* istanbul ignore next */\n
      labelPage(pageNumber) :
      \"\".concat(Object(__WEBPACK_IMPORTED_MODULE_8__utils_inspect__[\"f\" /*
      isFunction */])(labelPage) ? labelPage() : labelPage, \"
      \").concat(pageNumber),\n 'aria-checked': isNav ? null : active ? 'true' :
      'false',\n 'aria-current': isNav && active ? 'page' : null,\n
      'aria-posinset': isNav ? null : pageNumber,\n 'aria-setsize': isNav ? null
      : numberOfPages,\n // ARIA \"roving tabindex\" method (except in `isNav`
      mode)\n tabindex: isNav ? null : tabIndex\n };\n var btnContent =
      Object(__WEBPACK_IMPORTED_MODULE_15__utils_string__[\"g\" /* toString
      */])(_this7.makePage(pageNumber));\n var scope = {\n page: pageNumber,\n
      index: pageNumber - 1,\n content: btnContent,\n active: active,\n
      disabled: disabled\n };\n var $inner = h(disabled ? 'span' : isNav ?
      __WEBPACK_IMPORTED_MODULE_18__components_link_link__[\"a\" /* BLink */] :
      'button', {\n props: disabled || !isNav ? {} :
      _this7.linkProps(pageNumber),\n staticClass: 'page-link',\n class: {\n
      'flex-grow-1': !isNav && !disabled && fill\n },\n attrs: attrs,\n on:
      disabled ? {} : {\n '!click': function click(event) {\n
      _this7.onClick(event, pageNumber);\n },\n keydown: onSpaceKey\n }\n },
      [_this7.normalizeSlot(__WEBPACK_IMPORTED_MODULE_4__constants_slots__[\"W\"
      /* SLOT_NAME_PAGE */], scope) || btnContent]);\n return h('li', {\n
      staticClass: 'page-item',\n class: [{\n disabled: disabled,\n active:
      active,\n 'flex-fill': fill,\n 'd-flex': fill && !isNav && !disabled\n },
      page.classes, _this7.pageClass],\n attrs: {\n role: isNav ? null :
      'presentation'\n },\n key: \"page-\".concat(pageNumber)\n }, [$inner]);\n
      }; // Goto first page button\n // Don't render button when
      `hideGotoEndButtons` or `firstNumber` is set\n\n var $firstPageBtn =
      h();\n if (!this.firstNumber && !this.hideGotoEndButtons) {\n
      $firstPageBtn = makeEndBtn(1, this.labelFirstPage,
      __WEBPACK_IMPORTED_MODULE_4__constants_slots__[\"r\" /*
      SLOT_NAME_FIRST_TEXT */], this.firstText, this.firstClass, 1,
      'pagination-goto-first');\n }\n $buttons.push($firstPageBtn); // Goto
      previous page button\n\n $buttons.push(makeEndBtn(currentPage - 1,
      this.labelPrevPage, __WEBPACK_IMPORTED_MODULE_4__constants_slots__[\"Z\"
      /* SLOT_NAME_PREV_TEXT */], this.prevText, this.prevClass, 1,
      'pagination-goto-prev')); // Show first (1) button?\n\n
      $buttons.push(this.firstNumber && pageNumbers[0] !== 1 ?
      makePageButton({\n number: 1\n }, 0) : h()); // First ellipsis\n\n
      $buttons.push(showFirstDots ? makeEllipsis(false) : h()); // Individual
      page links\n\n this.pageList.forEach(function (page, idx) {\n var offset =
      showFirstDots && _this7.firstNumber && pageNumbers[0] !== 1 ? 1 : 0;\n
      $buttons.push(makePageButton(page, idx + offset));\n }); // Last
      ellipsis\n\n $buttons.push(showLastDots ? makeEllipsis(true) : h()); //
      Show last page button?\n\n $buttons.push(this.lastNumber &&
      pageNumbers[pageNumbers.length - 1] !== numberOfPages ? makePageButton({\n
      number: numberOfPages\n }, -1) : h()); // Goto next page button\n\n
      $buttons.push(makeEndBtn(currentPage + 1, this.labelNextPage,
      __WEBPACK_IMPORTED_MODULE_4__constants_slots__[\"U\" /*
      SLOT_NAME_NEXT_TEXT */], this.nextText, this.nextClass, numberOfPages,
      'pagination-goto-next')); // Goto last page button\n // Don't render
      button when `hideGotoEndButtons` or `lastNumber` is set\n\n var
      $lastPageBtn = h();\n if (!this.lastNumber && !this.hideGotoEndButtons)
      {\n $lastPageBtn = makeEndBtn(numberOfPages, this.labelLastPage,
      __WEBPACK_IMPORTED_MODULE_4__constants_slots__[\"D\" /*
      SLOT_NAME_LAST_TEXT */], this.lastText, this.lastClass, numberOfPages,
      'pagination-goto-last');\n }\n $buttons.push($lastPageBtn); // Assemble
      the pagination buttons\n\n var $pagination = h('ul', {\n staticClass:
      'pagination',\n class: ['b-pagination', this.btnSize, this.alignment,
      this.styleClass],\n attrs: {\n role: isNav ? null : 'menubar',\n
      'aria-disabled': disabled ? 'true' : 'false',\n 'aria-label': isNav ? null
      : ariaLabel || null\n },\n // We disable keyboard left/right nav when
      `<b-pagination-nav
        >`\n on: isNav ? {} : {\n keydown: this.handleKeyNav\n },\n ref: 'ul'\n
        }, $buttons); // If we are `<b-pagination-nav
          >`, wrap in `
          <nav>
            ` wrapper\n\n if (isNav) {\n return h('nav', {\n attrs: {\n
            'aria-disabled': disabled ? 'true' : null,\n 'aria-hidden': disabled
            ? 'true' : 'false',\n 'aria-label': isNav ? ariaLabel || null :
            null\n }\n }, [$pagination]);\n }\n return $pagination;\n }\n});//#
            sourceURL=[module]\n//#
            sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9ub2RlX21vZHVsZXMvYm9vdHN0cmFwLXZ1ZS9lc20vbWl4aW5zL3BhZ2luYXRpb24uanMuanMiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9ub2RlX21vZHVsZXMvYm9vdHN0cmFwLXZ1ZS9lc20vbWl4aW5zL3BhZ2luYXRpb24uanM/ZTc3NCJdLCJzb3VyY2VzQ29udGVudCI6WyJ2YXIgX3dhdGNoO1xuXG5mdW5jdGlvbiBvd25LZXlzKG9iamVjdCwgZW51bWVyYWJsZU9ubHkpIHsgdmFyIGtleXMgPSBPYmplY3Qua2V5cyhvYmplY3QpOyBpZiAoT2JqZWN0LmdldE93blByb3BlcnR5U3ltYm9scykgeyB2YXIgc3ltYm9scyA9IE9iamVjdC5nZXRPd25Qcm9wZXJ0eVN5bWJvbHMob2JqZWN0KTsgZW51bWVyYWJsZU9ubHkgJiYgKHN5bWJvbHMgPSBzeW1ib2xzLmZpbHRlcihmdW5jdGlvbiAoc3ltKSB7IHJldHVybiBPYmplY3QuZ2V0T3duUHJvcGVydHlEZXNjcmlwdG9yKG9iamVjdCwgc3ltKS5lbnVtZXJhYmxlOyB9KSksIGtleXMucHVzaC5hcHBseShrZXlzLCBzeW1ib2xzKTsgfSByZXR1cm4ga2V5czsgfVxuXG5mdW5jdGlvbiBfb2JqZWN0U3ByZWFkKHRhcmdldCkgeyBmb3IgKHZhciBpID0gMTsgaSA8IGFyZ3VtZW50cy5sZW5ndGg7IGkrKykgeyB2YXIgc291cmNlID0gbnVsbCAhPSBhcmd1bWVudHNbaV0gPyBhcmd1bWVudHNbaV0gOiB7fTsgaSAlIDIgPyBvd25LZXlzKE9iamVjdChzb3VyY2UpLCAhMCkuZm9yRWFjaChmdW5jdGlvbiAoa2V5KSB7IF9kZWZpbmVQcm9wZXJ0eSh0YXJnZXQsIGtleSwgc291cmNlW2tleV0pOyB9KSA6IE9iamVjdC5nZXRPd25Qcm9wZXJ0eURlc2NyaXB0b3JzID8gT2JqZWN0LmRlZmluZVByb3BlcnRpZXModGFyZ2V0LCBPYmplY3QuZ2V0T3duUHJvcGVydHlEZXNjcmlwdG9ycyhzb3VyY2UpKSA6IG93bktleXMoT2JqZWN0KHNvdXJjZSkpLmZvckVhY2goZnVuY3Rpb24gKGtleSkgeyBPYmplY3QuZGVmaW5lUHJvcGVydHkodGFyZ2V0LCBrZXksIE9iamVjdC5nZXRPd25Qcm9wZXJ0eURlc2NyaXB0b3Ioc291cmNlLCBrZXkpKTsgfSk7IH0gcmV0dXJuIHRhcmdldDsgfVxuXG5mdW5jdGlvbiBfZGVmaW5lUHJvcGVydHkob2JqLCBrZXksIHZhbHVlKSB7IGlmIChrZXkgaW4gb2JqKSB7IE9iamVjdC5kZWZpbmVQcm9wZXJ0eShvYmosIGtleSwgeyB2YWx1ZTogdmFsdWUsIGVudW1lcmFibGU6IHRydWUsIGNvbmZpZ3VyYWJsZTogdHJ1ZSwgd3JpdGFibGU6IHRydWUgfSk7IH0gZWxzZSB7IG9ialtrZXldID0gdmFsdWU7IH0gcmV0dXJuIG9iajsgfVxuXG5pbXBvcnQgeyBleHRlbmQgfSBmcm9tICcuLi92dWUnO1xuaW1wb3J0IHsgTkFNRV9QQUdJTkFUSU9OIH0gZnJvbSAnLi4vY29uc3RhbnRzL2NvbXBvbmVudHMnO1xuaW1wb3J0IHsgQ09ERV9ET1dOLCBDT0RFX0xFRlQsIENPREVfUklHSFQsIENPREVfU1BBQ0UsIENPREVfVVAgfSBmcm9tICcuLi9jb25zdGFudHMva2V5LWNvZGVzJztcbmltcG9ydCB7IFBST1BfVFlQRV9BUlJBWV9PQkpFQ1RfU1RSSU5HLCBQUk9QX1RZUEVfQk9PTEVBTiwgUFJPUF9UWVBFX0JPT0xFQU5fTlVNQkVSX1NUUklORywgUFJPUF9UWVBFX0ZVTkNUSU9OX1NUUklORywgUFJPUF9UWVBFX05VTUJFUl9TVFJJTkcsIFBST1BfVFlQRV9TVFJJTkcgfSBmcm9tICcuLi9jb25zdGFudHMvcHJvcHMnO1xuaW1wb3J0IHsgU0xPVF9OQU1FX0VMTElQU0lTX1RFWFQsIFNMT1RfTkFNRV9GSVJTVF9URVhULCBTTE9UX05BTUVfTEFTVF9URVhULCBTTE9UX05BTUVfTkVYVF9URVhULCBTTE9UX05BTUVfUEFHRSwgU0xPVF9OQU1FX1BSRVZfVEVYVCB9IGZyb20gJy4uL2NvbnN0YW50cy9zbG90cyc7XG5pbXBvcnQgeyBjcmVhdGVBcnJheSB9IGZyb20gJy4uL3V0aWxzL2FycmF5JztcbmltcG9ydCB7IGF0dGVtcHRGb2N1cywgZ2V0QWN0aXZlRWxlbWVudCwgZ2V0QXR0ciwgaXNEaXNhYmxlZCwgaXNWaXNpYmxlLCBzZWxlY3RBbGwgfSBmcm9tICcuLi91dGlscy9kb20nO1xuaW1wb3J0IHsgc3RvcEV2ZW50IH0gZnJvbSAnLi4vdXRpbHMvZXZlbnRzJztcbmltcG9ydCB7IGlzRnVuY3Rpb24sIGlzTnVsbCB9IGZyb20gJy4uL3V0aWxzL2luc3BlY3QnO1xuaW1wb3J0IHsgbWF0aEZsb29yLCBtYXRoTWF4LCBtYXRoTWluIH0gZnJvbSAnLi4vdXRpbHMvbWF0aCc7XG5pbXBvcnQgeyBtYWtlTW9kZWxNaXhpbiB9IGZyb20gJy4uL3V0aWxzL21vZGVsJztcbmltcG9ydCB7IHRvSW50ZWdlciB9IGZyb20gJy4uL3V0aWxzL251bWJlcic7XG5pbXBvcnQgeyBzb3J0S2V5cyB9IGZyb20gJy4uL3V0aWxzL29iamVjdCc7XG5pbXBvcnQgeyBoYXNQcm9wRnVuY3Rpb24sIG1ha2VQcm9wLCBtYWtlUHJvcHNDb25maWd1cmFibGUgfSBmcm9tICcuLi91dGlscy9wcm9wcyc7XG5pbXBvcnQgeyBzYWZlVnVlSW5zdGFuY2UgfSBmcm9tICcuLi91dGlscy9zYWZlLXZ1ZS1pbnN0YW5jZSc7XG5pbXBvcnQgeyB0b1N0cmluZyB9IGZyb20gJy4uL3V0aWxzL3N0cmluZyc7XG5pbXBvcnQgeyB3YXJuIH0gZnJvbSAnLi4vdXRpbHMvd2Fybic7XG5pbXBvcnQgeyBub3JtYWxpemVTbG90TWl4aW4gfSBmcm9tICcuLi9taXhpbnMvbm9ybWFsaXplLXNsb3QnO1xuaW1wb3J0IHsgQkxpbmsgfSBmcm9tICcuLi9jb21wb25lbnRzL2xpbmsvbGluayc7IC8vIENvbW1vbiBwcm9wcywgY29tcHV0ZWQsIGRhdGEsIHJlbmRlciBmdW5jdGlvbiwgYW5kIG1ldGhvZHNcbi8vIGZvciBgPGItcGFnaW5hdGlvbj5gIGFuZCBgPGItcGFnaW5hdGlvbi1uYXY+YFxuLy8gLS0tIENvbnN0YW50cyAtLS1cblxudmFyIF9tYWtlTW9kZWxNaXhpbiA9IG1ha2VNb2RlbE1peGluKCd2YWx1ZScsIHtcbiAgdHlwZTogUFJPUF9UWVBFX0JPT0xFQU5fTlVNQkVSX1NUUklORyxcbiAgZGVmYXVsdFZhbHVlOiBudWxsLFxuXG4gIC8qIGlzdGFuYnVsIGlnbm9yZSBuZXh0ICovXG4gIHZhbGlkYXRvcjogZnVuY3Rpb24gdmFsaWRhdG9yKHZhbHVlKSB7XG4gICAgaWYgKCFpc051bGwodmFsdWUpICYmIHRvSW50ZWdlcih2YWx1ZSwgMCkgPCAxKSB7XG4gICAgICB3YXJuKCdcInYtbW9kZWxcIiB2YWx1ZSBtdXN0IGJlIGEgbnVtYmVyIGdyZWF0ZXIgdGhhbiBcIjBcIicsIE5BTUVfUEFHSU5BVElPTik7XG4gICAgICByZXR1cm4gZmFsc2U7XG4gICAgfVxuXG4gICAgcmV0dXJuIHRydWU7XG4gIH1cbn0pLFxuICAgIG1vZGVsTWl4aW4gPSBfbWFrZU1vZGVsTWl4aW4ubWl4aW4sXG4gICAgbW9kZWxQcm9wcyA9IF9tYWtlTW9kZWxNaXhpbi5wcm9wcyxcbiAgICBNT0RFTF9QUk9QX05BTUUgPSBfbWFrZU1vZGVsTWl4aW4ucHJvcCxcbiAgICBNT0RFTF9FVkVOVF9OQU1FID0gX21ha2VNb2RlbE1peGluLmV2ZW50O1xuXG5leHBvcnQgeyBNT0RFTF9QUk9QX05BTUUsIE1PREVMX0VWRU5UX05BTUUgfTsgLy8gVGhyZXNob2xkIG9mIGxpbWl0IHNpemUgd2hlbiB3ZSBzdGFydC9zdG9wIHNob3dpbmcgZWxsaXBzaXNcblxudmFyIEVMTElQU0lTX1RIUkVTSE9MRCA9IDM7IC8vIERlZmF1bHQgIyBvZiBidXR0b25zIGxpbWl0XG5cbnZhciBERUZBVUxUX0xJTUlUID0gNTsgLy8gLS0tIEhlbHBlciBtZXRob2RzIC0tLVxuLy8gTWFrZSBhbiBhcnJheSBvZiBOIHRvIE4rWFxuXG52YXIgbWFrZVBhZ2VBcnJheSA9IGZ1bmN0aW9uIG1ha2VQYWdlQXJyYXkoc3RhcnROdW1iZXIsIG51bWJlck9mUGFnZXMpIHtcbiAgcmV0dXJuIGNyZWF0ZUFycmF5KG51bWJlck9mUGFnZXMsIGZ1bmN0aW9uIChfLCBpKSB7XG4gICAgcmV0dXJuIHtcbiAgICAgIG51bWJlcjogc3RhcnROdW1iZXIgKyBpLFxuICAgICAgY2xhc3NlczogbnVsbFxuICAgIH07XG4gIH0pO1xufTsgLy8gU2FuaXRpemUgdGhlIHByb3ZpZGVkIGxpbWl0IHZhbHVlIChjb252ZXJ0aW5nIHRvIGEgbnVtYmVyKVxuXG5cbnZhciBzYW5pdGl6ZUxpbWl0ID0gZnVuY3Rpb24gc2FuaXRpemVMaW1pdCh2YWx1ZSkge1xuICB2YXIgbGltaXQgPSB0b0ludGVnZXIodmFsdWUpIHx8IDE7XG4gIHJldHVybiBsaW1pdCA8IDEgPyBERUZBVUxUX0xJTUlUIDogbGltaXQ7XG59OyAvLyBTYW5pdGl6ZSB0aGUgcHJvdmlkZWQgY3VycmVudCBwYWdlIG51bWJlciAoY29udmVydGluZyB0byBhIG51bWJlcilcblxuXG52YXIgc2FuaXRpemVDdXJyZW50UGFnZSA9IGZ1bmN0aW9uIHNhbml0aXplQ3VycmVudFBhZ2UodmFsLCBudW1iZXJPZlBhZ2VzKSB7XG4gIHZhciBwYWdlID0gdG9JbnRlZ2VyKHZhbCkgfHwgMTtcbiAgcmV0dXJuIHBhZ2UgPiBudW1iZXJPZlBhZ2VzID8gbnVtYmVyT2ZQYWdlcyA6IHBhZ2UgPCAxID8gMSA6IHBhZ2U7XG59OyAvLyBMaW5rcyBkb24ndCBub3JtYWxseSByZXNwb25kIHRvIFNQQUNFLCBzbyB3ZSBhZGQgdGhhdFxuLy8gZnVuY3Rpb25hbGl0eSB2aWEgdGhpcyBoYW5kbGVyXG5cblxudmFyIG9uU3BhY2VLZXkgPSBmdW5jdGlvbiBvblNwYWNlS2V5KGV2ZW50KSB7XG4gIGlmIChldmVudC5rZXlDb2RlID09PSBDT0RFX1NQQUNFKSB7XG4gICAgLy8gU3RvcCBwYWdlIGZyb20gc2Nyb2xsaW5nXG4gICAgc3RvcEV2ZW50KGV2ZW50LCB7XG4gICAgICBpbW1lZGlhdGVQcm9wYWdhdGlvbjogdHJ1ZVxuICAgIH0pOyAvLyBUcmlnZ2VyIHRoZSBjbGljayBldmVudCBvbiB0aGUgbGlua1xuXG4gICAgZXZlbnQuY3VycmVudFRhcmdldC5jbGljaygpO1xuICAgIHJldHVybiBmYWxzZTtcbiAgfVxufTsgLy8gLS0tIFByb3BzIC0tLVxuXG5cbmV4cG9ydCB2YXIgcHJvcHMgPSBtYWtlUHJvcHNDb25maWd1cmFibGUoc29ydEtleXMoX29iamVjdFNwcmVhZChfb2JqZWN0U3ByZWFkKHt9LCBtb2RlbFByb3BzKSwge30sIHtcbiAgYWxpZ246IG1ha2VQcm9wKFBST1BfVFlQRV9TVFJJTkcsICdsZWZ0JyksXG4gIGFyaWFMYWJlbDogbWFrZVByb3AoUFJPUF9UWVBFX1NUUklORywgJ1BhZ2luYXRpb24nKSxcbiAgZGlzYWJsZWQ6IG1ha2VQcm9wKFBST1BfVFlQRV9CT09MRUFOLCBmYWxzZSksXG4gIGVsbGlwc2lzQ2xhc3M6IG1ha2VQcm9wKFBST1BfVFlQRV9BUlJBWV9PQkpFQ1RfU1RSSU5HKSxcbiAgZWxsaXBzaXNUZXh0OiBtYWtlUHJvcChQUk9QX1RZUEVfU1RSSU5HLCBcIlxcdTIwMjZcIiksXG4gIC8vICfigKYnXG4gIGZpcnN0Q2xhc3M6IG1ha2VQcm9wKFBST1BfVFlQRV9BUlJBWV9PQkpFQ1RfU1RSSU5HKSxcbiAgZmlyc3ROdW1iZXI6IG1ha2VQcm9wKFBST1BfVFlQRV9CT09MRUFOLCBmYWxzZSksXG4gIGZpcnN0VGV4dDogbWFrZVByb3AoUFJPUF9UWVBFX1NUUklORywgXCJcXHhBQlwiKSxcbiAgLy8gJ8KrJ1xuICBoaWRlRWxsaXBzaXM6IG1ha2VQcm9wKFBST1BfVFlQRV9CT09MRUFOLCBmYWxzZSksXG4gIGhpZGVHb3RvRW5kQnV0dG9uczogbWFrZVByb3AoUFJPUF9UWVBFX0JPT0xFQU4sIGZhbHNlKSxcbiAgbGFiZWxGaXJzdFBhZ2U6IG1ha2VQcm9wKFBST1BfVFlQRV9TVFJJTkcsICdHbyB0byBmaXJzdCBwYWdlJyksXG4gIGxhYmVsTGFzdFBhZ2U6IG1ha2VQcm9wKFBST1BfVFlQRV9TVFJJTkcsICdHbyB0byBsYXN0IHBhZ2UnKSxcbiAgbGFiZWxOZXh0UGFnZTogbWFrZVByb3AoUFJPUF9UWVBFX1NUUklORywgJ0dvIHRvIG5leHQgcGFnZScpLFxuICBsYWJlbFBhZ2U6IG1ha2VQcm9wKFBST1BfVFlQRV9GVU5DVElPTl9TVFJJTkcsICdHbyB0byBwYWdlJyksXG4gIGxhYmVsUHJldlBhZ2U6IG1ha2VQcm9wKFBST1BfVFlQRV9TVFJJTkcsICdHbyB0byBwcmV2aW91cyBwYWdlJyksXG4gIGxhc3RDbGFzczogbWFrZVByb3AoUFJPUF9UWVBFX0FSUkFZX09CSkVDVF9TVFJJTkcpLFxuICBsYXN0TnVtYmVyOiBtYWtlUHJvcChQUk9QX1RZUEVfQk9PTEVBTiwgZmFsc2UpLFxuICBsYXN0VGV4dDogbWFrZVByb3AoUFJPUF9UWVBFX1NUUklORywgXCJcXHhCQlwiKSxcbiAgLy8gJ8K7J1xuICBsaW1pdDogbWFrZVByb3AoUFJPUF9UWVBFX05VTUJFUl9TVFJJTkcsIERFRkFVTFRfTElNSVQsXG4gIC8qIGlzdGFuYnVsIGlnbm9yZSBuZXh0ICovXG4gIGZ1bmN0aW9uICh2YWx1ZSkge1xuICAgIGlmICh0b0ludGVnZXIodmFsdWUsIDApIDwgMSkge1xuICAgICAgd2FybignUHJvcCBcImxpbWl0XCIgbXVzdCBiZSBhIG51bWJlciBncmVhdGVyIHRoYW4gXCIwXCInLCBOQU1FX1BBR0lOQVRJT04pO1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cblxuICAgIHJldHVybiB0cnVlO1xuICB9KSxcbiAgbmV4dENsYXNzOiBtYWtlUHJvcChQUk9QX1RZUEVfQVJSQVlfT0JKRUNUX1NUUklORyksXG4gIG5leHRUZXh0OiBtYWtlUHJvcChQUk9QX1RZUEVfU1RSSU5HLCBcIlxcdTIwM0FcIiksXG4gIC8vICfigLonXG4gIHBhZ2VDbGFzczogbWFrZVByb3AoUFJPUF9UWVBFX0FSUkFZX09CSkVDVF9TVFJJTkcpLFxuICBwaWxsczogbWFrZVByb3AoUFJPUF9UWVBFX0JPT0xFQU4sIGZhbHNlKSxcbiAgcHJldkNsYXNzOiBtYWtlUHJvcChQUk9QX1RZUEVfQVJSQVlfT0JKRUNUX1NUUklORyksXG4gIHByZXZUZXh0OiBtYWtlUHJvcChQUk9QX1RZUEVfU1RSSU5HLCBcIlxcdTIwMzlcIiksXG4gIC8vICfigLknXG4gIHNpemU6IG1ha2VQcm9wKFBST1BfVFlQRV9TVFJJTkcpXG59KSksICdwYWdpbmF0aW9uJyk7IC8vIC0tLSBNaXhpbiAtLS1cbi8vIEB2dWUvY29tcG9uZW50XG5cbmV4cG9ydCB2YXIgcGFnaW5hdGlvbk1peGluID0gZXh0ZW5kKHtcbiAgbWl4aW5zOiBbbW9kZWxNaXhpbiwgbm9ybWFsaXplU2xvdE1peGluXSxcbiAgcHJvcHM6IHByb3BzLFxuICBkYXRhOiBmdW5jdGlvbiBkYXRhKCkge1xuICAgIC8vIGAtMWAgc2lnbmlmaWVzIG5vIHBhZ2UgaW5pdGlhbGx5IHNlbGVjdGVkXG4gICAgdmFyIGN1cnJlbnRQYWdlID0gdG9JbnRlZ2VyKHRoaXNbTU9ERUxfUFJPUF9OQU1FXSwgMCk7XG4gICAgY3VycmVudFBhZ2UgPSBjdXJyZW50UGFnZSA+IDAgPyBjdXJyZW50UGFnZSA6IC0xO1xuICAgIHJldHVybiB7XG4gICAgICBjdXJyZW50UGFnZTogY3VycmVudFBhZ2UsXG4gICAgICBsb2NhbE51bWJlck9mUGFnZXM6IDEsXG4gICAgICBsb2NhbExpbWl0OiBERUZBVUxUX0xJTUlUXG4gICAgfTtcbiAgfSxcbiAgY29tcHV0ZWQ6IHtcbiAgICBidG5TaXplOiBmdW5jdGlvbiBidG5TaXplKCkge1xuICAgICAgdmFyIHNpemUgPSB0aGlzLnNpemU7XG4gICAgICByZXR1cm4gc2l6ZSA/IFwicGFnaW5hdGlvbi1cIi5jb25jYXQoc2l6ZSkgOiAnJztcbiAgICB9LFxuICAgIGFsaWdubWVudDogZnVuY3Rpb24gYWxpZ25tZW50KCkge1xuICAgICAgdmFyIGFsaWduID0gdGhpcy5hbGlnbjtcblxuICAgICAgaWYgKGFsaWduID09PSAnY2VudGVyJykge1xuICAgICAgICByZXR1cm4gJ2p1c3RpZnktY29udGVudC1jZW50ZXInO1xuICAgICAgfSBlbHNlIGlmIChhbGlnbiA9PT0gJ2VuZCcgfHwgYWxpZ24gPT09ICdyaWdodCcpIHtcbiAgICAgICAgcmV0dXJuICdqdXN0aWZ5LWNvbnRlbnQtZW5kJztcbiAgICAgIH0gZWxzZSBpZiAoYWxpZ24gPT09ICdmaWxsJykge1xuICAgICAgICAvLyBUaGUgcGFnZS1pdGVtcyB3aWxsIGFsc28gaGF2ZSAnZmxleC1maWxsJyBhZGRlZFxuICAgICAgICAvLyBXZSBhZGQgdGV4dCBjZW50ZXJpbmcgdG8gbWFrZSB0aGUgYnV0dG9uIGFwcGVhcmFuY2UgYmV0dGVyIGluIGZpbGwgbW9kZVxuICAgICAgICByZXR1cm4gJ3RleHQtY2VudGVyJztcbiAgICAgIH1cblxuICAgICAgcmV0dXJuICcnO1xuICAgIH0sXG4gICAgc3R5bGVDbGFzczogZnVuY3Rpb24gc3R5bGVDbGFzcygpIHtcbiAgICAgIHJldHVybiB0aGlzLnBpbGxzID8gJ2ItcGFnaW5hdGlvbi1waWxscycgOiAnJztcbiAgICB9LFxuICAgIGNvbXB1dGVkQ3VycmVudFBhZ2U6IGZ1bmN0aW9uIGNvbXB1dGVkQ3VycmVudFBhZ2UoKSB7XG4gICAgICByZXR1cm4gc2FuaXRpemVDdXJyZW50UGFnZSh0aGlzLmN1cnJlbnRQYWdlLCB0aGlzLmxvY2FsTnVtYmVyT2ZQYWdlcyk7XG4gICAgfSxcbiAgICBwYWdpbmF0aW9uUGFyYW1zOiBmdW5jdGlvbiBwYWdpbmF0aW9uUGFyYW1zKCkge1xuICAgICAgLy8gRGV0ZXJtaW5lIGlmIHdlIHNob3VsZCBzaG93IHRoZSB0aGUgZWxsaXBzaXNcbiAgICAgIHZhciBsaW1pdCA9IHRoaXMubG9jYWxMaW1pdCxcbiAgICAgICAgICBudW1iZXJPZlBhZ2VzID0gdGhpcy5sb2NhbE51bWJlck9mUGFnZXMsXG4gICAgICAgICAgY3VycmVudFBhZ2UgPSB0aGlzLmNvbXB1dGVkQ3VycmVudFBhZ2UsXG4gICAgICAgICAgaGlkZUVsbGlwc2lzID0gdGhpcy5oaWRlRWxsaXBzaXMsXG4gICAgICAgICAgZmlyc3ROdW1iZXIgPSB0aGlzLmZpcnN0TnVtYmVyLFxuICAgICAgICAgIGxhc3ROdW1iZXIgPSB0aGlzLmxhc3ROdW1iZXI7XG4gICAgICB2YXIgc2hvd0ZpcnN0RG90cyA9IGZhbHNlO1xuICAgICAgdmFyIHNob3dMYXN0RG90cyA9IGZhbHNlO1xuICAgICAgdmFyIG51bWJlck9mTGlua3MgPSBsaW1pdDtcbiAgICAgIHZhciBzdGFydE51bWJlciA9IDE7XG5cbiAgICAgIGlmIChudW1iZXJPZlBhZ2VzIDw9IGxpbWl0KSB7XG4gICAgICAgIC8vIFNwZWNpYWwgY2FzZTogTGVzcyBwYWdlcyBhdmFpbGFibGUgdGhhbiB0aGUgbGltaXQgb2YgZGlzcGxheWVkIHBhZ2VzXG4gICAgICAgIG51bWJlck9mTGlua3MgPSBudW1iZXJPZlBhZ2VzO1xuICAgICAgfSBlbHNlIGlmIChjdXJyZW50UGFnZSA8IGxpbWl0IC0gMSAmJiBsaW1pdCA+IEVMTElQU0lTX1RIUkVTSE9MRCkge1xuICAgICAgICBpZiAoIWhpZGVFbGxpcHNpcyB8fCBsYXN0TnVtYmVyKSB7XG4gICAgICAgICAgc2hvd0xhc3REb3RzID0gdHJ1ZTtcbiAgICAgICAgICBudW1iZXJPZkxpbmtzID0gbGltaXQgLSAoZmlyc3ROdW1iZXIgPyAwIDogMSk7XG4gICAgICAgIH1cblxuICAgICAgICBudW1iZXJPZkxpbmtzID0gbWF0aE1pbihudW1iZXJPZkxpbmtzLCBsaW1pdCk7XG4gICAgICB9IGVsc2UgaWYgKG51bWJlck9mUGFnZXMgLSBjdXJyZW50UGFnZSArIDIgPCBsaW1pdCAmJiBsaW1pdCA+IEVMTElQU0lTX1RIUkVTSE9MRCkge1xuICAgICAgICBpZiAoIWhpZGVFbGxpcHNpcyB8fCBmaXJzdE51bWJlcikge1xuICAgICAgICAgIHNob3dGaXJzdERvdHMgPSB0cnVlO1xuICAgICAgICAgIG51bWJlck9mTGlua3MgPSBsaW1pdCAtIChsYXN0TnVtYmVyID8gMCA6IDEpO1xuICAgICAgICB9XG5cbiAgICAgICAgc3RhcnROdW1iZXIgPSBudW1iZXJPZlBhZ2VzIC0gbnVtYmVyT2ZMaW5rcyArIDE7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICAvLyBXZSBhcmUgc29tZXdoZXJlIGluIHRoZSBtaWRkbGUgb2YgdGhlIHBhZ2UgbGlzdFxuICAgICAgICBpZiAobGltaXQgPiBFTExJUFNJU19USFJFU0hPTEQpIHtcbiAgICAgICAgICBudW1iZXJPZkxpbmtzID0gbGltaXQgLSAoaGlkZUVsbGlwc2lzID8gMCA6IDIpO1xuICAgICAgICAgIHNob3dGaXJzdERvdHMgPSAhISghaGlkZUVsbGlwc2lzIHx8IGZpcnN0TnVtYmVyKTtcbiAgICAgICAgICBzaG93TGFzdERvdHMgPSAhISghaGlkZUVsbGlwc2lzIHx8IGxhc3ROdW1iZXIpO1xuICAgICAgICB9XG5cbiAgICAgICAgc3RhcnROdW1iZXIgPSBjdXJyZW50UGFnZSAtIG1hdGhGbG9vcihudW1iZXJPZkxpbmtzIC8gMik7XG4gICAgICB9IC8vIFNhbml0eSBjaGVja3NcblxuICAgICAgLyogaXN0YW5idWwgaWdub3JlIGlmICovXG5cblxuICAgICAgaWYgKHN0YXJ0TnVtYmVyIDwgMSkge1xuICAgICAgICBzdGFydE51bWJlciA9IDE7XG4gICAgICAgIHNob3dGaXJzdERvdHMgPSBmYWxzZTtcbiAgICAgIH0gZWxzZSBpZiAoc3RhcnROdW1iZXIgPiBudW1iZXJPZlBhZ2VzIC0gbnVtYmVyT2ZMaW5rcykge1xuICAgICAgICBzdGFydE51bWJlciA9IG51bWJlck9mUGFnZXMgLSBudW1iZXJPZkxpbmtzICsgMTtcbiAgICAgICAgc2hvd0xhc3REb3RzID0gZmFsc2U7XG4gICAgICB9XG5cbiAgICAgIGlmIChzaG93Rmlyc3REb3RzICYmIGZpcnN0TnVtYmVyICYmIHN0YXJ0TnVtYmVyIDwgNCkge1xuICAgICAgICBudW1iZXJPZkxpbmtzID0gbnVtYmVyT2ZMaW5rcyArIDI7XG4gICAgICAgIHN0YXJ0TnVtYmVyID0gMTtcbiAgICAgICAgc2hvd0ZpcnN0RG90cyA9IGZhbHNlO1xuICAgICAgfVxuXG4gICAgICB2YXIgbGFzdFBhZ2VOdW1iZXIgPSBzdGFydE51bWJlciArIG51bWJlck9mTGlua3MgLSAxO1xuXG4gICAgICBpZiAoc2hvd0xhc3REb3RzICYmIGxhc3ROdW1iZXIgJiYgbGFzdFBhZ2VOdW1iZXIgPiBudW1iZXJPZlBhZ2VzIC0gMykge1xuICAgICAgICBudW1iZXJPZkxpbmtzID0gbnVtYmVyT2ZMaW5rcyArIChsYXN0UGFnZU51bWJlciA9PT0gbnVtYmVyT2ZQYWdlcyAtIDIgPyAyIDogMyk7XG4gICAgICAgIHNob3dMYXN0RG90cyA9IGZhbHNlO1xuICAgICAgfSAvLyBTcGVjaWFsIGhhbmRsaW5nIGZvciBsb3dlciBsaW1pdHMgKHdoZXJlIGVsbGlwc2lzIGFyZSBuZXZlciBzaG93bilcblxuXG4gICAgICBpZiAobGltaXQgPD0gRUxMSVBTSVNfVEhSRVNIT0xEKSB7XG4gICAgICAgIGlmIChmaXJzdE51bWJlciAmJiBzdGFydE51bWJlciA9PT0gMSkge1xuICAgICAgICAgIG51bWJlck9mTGlua3MgPSBtYXRoTWluKG51bWJlck9mTGlua3MgKyAxLCBudW1iZXJPZlBhZ2VzLCBsaW1pdCArIDEpO1xuICAgICAgICB9IGVsc2UgaWYgKGxhc3ROdW1iZXIgJiYgbnVtYmVyT2ZQYWdlcyA9PT0gc3RhcnROdW1iZXIgKyBudW1iZXJPZkxpbmtzIC0gMSkge1xuICAgICAgICAgIHN0YXJ0TnVtYmVyID0gbWF0aE1heChzdGFydE51bWJlciAtIDEsIDEpO1xuICAgICAgICAgIG51bWJlck9mTGlua3MgPSBtYXRoTWluKG51bWJlck9mUGFnZXMgLSBzdGFydE51bWJlciArIDEsIG51bWJlck9mUGFnZXMsIGxpbWl0ICsgMSk7XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgbnVtYmVyT2ZMaW5rcyA9IG1hdGhNaW4obnVtYmVyT2ZMaW5rcywgbnVtYmVyT2ZQYWdlcyAtIHN0YXJ0TnVtYmVyICsgMSk7XG4gICAgICByZXR1cm4ge1xuICAgICAgICBzaG93Rmlyc3REb3RzOiBzaG93Rmlyc3REb3RzLFxuICAgICAgICBzaG93TGFzdERvdHM6IHNob3dMYXN0RG90cyxcbiAgICAgICAgbnVtYmVyT2ZMaW5rczogbnVtYmVyT2ZMaW5rcyxcbiAgICAgICAgc3RhcnROdW1iZXI6IHN0YXJ0TnVtYmVyXG4gICAgICB9O1xuICAgIH0sXG4gICAgcGFnZUxpc3Q6IGZ1bmN0aW9uIHBhZ2VMaXN0KCkge1xuICAgICAgLy8gR2VuZXJhdGVzIHRoZSBwYWdlTGlzdCBhcnJheVxuICAgICAgdmFyIF90aGlzJHBhZ2luYXRpb25QYXJhbSA9IHRoaXMucGFnaW5hdGlvblBhcmFtcyxcbiAgICAgICAgICBudW1iZXJPZkxpbmtzID0gX3RoaXMkcGFnaW5hdGlvblBhcmFtLm51bWJlck9mTGlua3MsXG4gICAgICAgICAgc3RhcnROdW1iZXIgPSBfdGhpcyRwYWdpbmF0aW9uUGFyYW0uc3RhcnROdW1iZXI7XG4gICAgICB2YXIgY3VycmVudFBhZ2UgPSB0aGlzLmNvbXB1dGVkQ3VycmVudFBhZ2U7IC8vIEdlbmVyYXRlIGxpc3Qgb2YgcGFnZSBudW1iZXJzXG5cbiAgICAgIHZhciBwYWdlcyA9IG1ha2VQYWdlQXJyYXkoc3RhcnROdW1iZXIsIG51bWJlck9mTGlua3MpOyAvLyBXZSBsaW1pdCB0byBhIHRvdGFsIG9mIDMgcGFnZSBidXR0b25zIG9uIFhTIHNjcmVlbnNcbiAgICAgIC8vIFNvIGFkZCBjbGFzc2VzIHRvIHBhZ2UgbGlua3MgdG8gaGlkZSB0aGVtIGZvciBYUyBicmVha3BvaW50XG4gICAgICAvLyBOb3RlOiBFbGxpcHNpcyB3aWxsIGFsc28gYmUgaGlkZGVuIG9uIFhTIHNjcmVlbnNcbiAgICAgIC8vIFRPRE86IE1ha2UgdGhpcyB2aXN1YWwgbGltaXQgY29uZmlndXJhYmxlIGJhc2VkIG9uIGJyZWFrcG9pbnQocylcblxuICAgICAgaWYgKHBhZ2VzLmxlbmd0aCA+IDMpIHtcbiAgICAgICAgdmFyIGlkeCA9IGN1cnJlbnRQYWdlIC0gc3RhcnROdW1iZXI7IC8vIFRIZSBmb2xsb3dpbmcgaXMgYSBib290c3RyYXAtdnVlIGN1c3RvbSB1dGlsaXR5IGNsYXNzXG5cbiAgICAgICAgdmFyIGNsYXNzZXMgPSAnYnYtZC14cy1kb3duLW5vbmUnO1xuXG4gICAgICAgIGlmIChpZHggPT09IDApIHtcbiAgICAgICAgICAvLyBLZWVwIGxlZnRtb3N0IDMgYnV0dG9ucyB2aXNpYmxlIHdoZW4gY3VycmVudCBwYWdlIGlzIGZpcnN0IHBhZ2VcbiAgICAgICAgICBmb3IgKHZhciBpID0gMzsgaSA8IHBhZ2VzLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgICAgICBwYWdlc1tpXS5jbGFzc2VzID0gY2xhc3NlcztcbiAgICAgICAgICB9XG4gICAgICAgIH0gZWxzZSBpZiAoaWR4ID09PSBwYWdlcy5sZW5ndGggLSAxKSB7XG4gICAgICAgICAgLy8gS2VlcCByaWdodG1vc3QgMyBidXR0b25zIHZpc2libGUgd2hlbiBjdXJyZW50IHBhZ2UgaXMgbGFzdCBwYWdlXG4gICAgICAgICAgZm9yICh2YXIgX2kgPSAwOyBfaSA8IHBhZ2VzLmxlbmd0aCAtIDM7IF9pKyspIHtcbiAgICAgICAgICAgIHBhZ2VzW19pXS5jbGFzc2VzID0gY2xhc3NlcztcbiAgICAgICAgICB9XG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgLy8gSGlkZSBhbGwgZXhjZXB0IGN1cnJlbnQgcGFnZSwgY3VycmVudCBwYWdlIC0gMSBhbmQgY3VycmVudCBwYWdlICsgMVxuICAgICAgICAgIGZvciAodmFyIF9pMiA9IDA7IF9pMiA8IGlkeCAtIDE7IF9pMisrKSB7XG4gICAgICAgICAgICAvLyBoaWRlIHNvbWUgbGVmdCBidXR0b24ocylcbiAgICAgICAgICAgIHBhZ2VzW19pMl0uY2xhc3NlcyA9IGNsYXNzZXM7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgZm9yICh2YXIgX2kzID0gcGFnZXMubGVuZ3RoIC0gMTsgX2kzID4gaWR4ICsgMTsgX2kzLS0pIHtcbiAgICAgICAgICAgIC8vIGhpZGUgc29tZSByaWdodCBidXR0b24ocylcbiAgICAgICAgICAgIHBhZ2VzW19pM10uY2xhc3NlcyA9IGNsYXNzZXM7XG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9XG5cbiAgICAgIHJldHVybiBwYWdlcztcbiAgICB9XG4gIH0sXG4gIHdhdGNoOiAoX3dhdGNoID0ge30sIF9kZWZpbmVQcm9wZXJ0eShfd2F0Y2gsIE1PREVMX1BST1BfTkFNRSwgZnVuY3Rpb24gKG5ld1ZhbHVlLCBvbGRWYWx1ZSkge1xuICAgIGlmIChuZXdWYWx1ZSAhPT0gb2xkVmFsdWUpIHtcbiAgICAgIHRoaXMuY3VycmVudFBhZ2UgPSBzYW5pdGl6ZUN1cnJlbnRQYWdlKG5ld1ZhbHVlLCB0aGlzLmxvY2FsTnVtYmVyT2ZQYWdlcyk7XG4gICAgfVxuICB9KSwgX2RlZmluZVByb3BlcnR5KF93YXRjaCwgXCJjdXJyZW50UGFnZVwiLCBmdW5jdGlvbiBjdXJyZW50UGFnZShuZXdWYWx1ZSwgb2xkVmFsdWUpIHtcbiAgICBpZiAobmV3VmFsdWUgIT09IG9sZFZhbHVlKSB7XG4gICAgICAvLyBFbWl0IGBudWxsYCBpZiBubyBwYWdlIHNlbGVjdGVkXG4gICAgICB0aGlzLiRlbWl0KE1PREVMX0VWRU5UX05BTUUsIG5ld1ZhbHVlID4gMCA/IG5ld1ZhbHVlIDogbnVsbCk7XG4gICAgfVxuICB9KSwgX2RlZmluZVByb3BlcnR5KF93YXRjaCwgXCJsaW1pdFwiLCBmdW5jdGlvbiBsaW1pdChuZXdWYWx1ZSwgb2xkVmFsdWUpIHtcbiAgICBpZiAobmV3VmFsdWUgIT09IG9sZFZhbHVlKSB7XG4gICAgICB0aGlzLmxvY2FsTGltaXQgPSBzYW5pdGl6ZUxpbWl0KG5ld1ZhbHVlKTtcbiAgICB9XG4gIH0pLCBfd2F0Y2gpLFxuICBjcmVhdGVkOiBmdW5jdGlvbiBjcmVhdGVkKCkge1xuICAgIHZhciBfdGhpcyA9IHRoaXM7XG5cbiAgICAvLyBTZXQgb3VyIGRlZmF1bHQgdmFsdWVzIGluIGRhdGFcbiAgICB0aGlzLmxvY2FsTGltaXQgPSBzYW5pdGl6ZUxpbWl0KHRoaXMubGltaXQpO1xuICAgIHRoaXMuJG5leHRUaWNrKGZ1bmN0aW9uICgpIHtcbiAgICAgIC8vIFNhbml0eSBjaGVja1xuICAgICAgX3RoaXMuY3VycmVudFBhZ2UgPSBfdGhpcy5jdXJyZW50UGFnZSA+IF90aGlzLmxvY2FsTnVtYmVyT2ZQYWdlcyA/IF90aGlzLmxvY2FsTnVtYmVyT2ZQYWdlcyA6IF90aGlzLmN1cnJlbnRQYWdlO1xuICAgIH0pO1xuICB9LFxuICBtZXRob2RzOiB7XG4gICAgaGFuZGxlS2V5TmF2OiBmdW5jdGlvbiBoYW5kbGVLZXlOYXYoZXZlbnQpIHtcbiAgICAgIHZhciBrZXlDb2RlID0gZXZlbnQua2V5Q29kZSxcbiAgICAgICAgICBzaGlmdEtleSA9IGV2ZW50LnNoaWZ0S2V5O1xuICAgICAgLyogaXN0YW5idWwgaWdub3JlIGlmICovXG5cbiAgICAgIGlmICh0aGlzLmlzTmF2KSB7XG4gICAgICAgIC8vIFdlIGRpc2FibGUgbGVmdC9yaWdodCBrZXlib2FyZCBuYXZpZ2F0aW9uIGluIGA8Yi1wYWdpbmF0aW9uLW5hdj5gXG4gICAgICAgIHJldHVybjtcbiAgICAgIH1cblxuICAgICAgaWYgKGtleUNvZGUgPT09IENPREVfTEVGVCB8fCBrZXlDb2RlID09PSBDT0RFX1VQKSB7XG4gICAgICAgIHN0b3BFdmVudChldmVudCwge1xuICAgICAgICAgIHByb3BhZ2F0aW9uOiBmYWxzZVxuICAgICAgICB9KTtcbiAgICAgICAgc2hpZnRLZXkgPyB0aGlzLmZvY3VzRmlyc3QoKSA6IHRoaXMuZm9jdXNQcmV2KCk7XG4gICAgICB9IGVsc2UgaWYgKGtleUNvZGUgPT09IENPREVfUklHSFQgfHwga2V5Q29kZSA9PT0gQ09ERV9ET1dOKSB7XG4gICAgICAgIHN0b3BFdmVudChldmVudCwge1xuICAgICAgICAgIHByb3BhZ2F0aW9uOiBmYWxzZVxuICAgICAgICB9KTtcbiAgICAgICAgc2hpZnRLZXkgPyB0aGlzLmZvY3VzTGFzdCgpIDogdGhpcy5mb2N1c05leHQoKTtcbiAgICAgIH1cbiAgICB9LFxuICAgIGdldEJ1dHRvbnM6IGZ1bmN0aW9uIGdldEJ1dHRvbnMoKSB7XG4gICAgICAvLyBSZXR1cm4gb25seSBidXR0b25zIHRoYXQgYXJlIHZpc2libGVcbiAgICAgIHJldHVybiBzZWxlY3RBbGwoJ2J1dHRvbi5wYWdlLWxpbmssIGEucGFnZS1saW5rJywgdGhpcy4kZWwpLmZpbHRlcihmdW5jdGlvbiAoYnRuKSB7XG4gICAgICAgIHJldHVybiBpc1Zpc2libGUoYnRuKTtcbiAgICAgIH0pO1xuICAgIH0sXG4gICAgZm9jdXNDdXJyZW50OiBmdW5jdGlvbiBmb2N1c0N1cnJlbnQoKSB7XG4gICAgICB2YXIgX3RoaXMyID0gdGhpcztcblxuICAgICAgLy8gV2UgZG8gdGhpcyBpbiBgJG5leHRUaWNrKClgIHRvIGVuc3VyZSBidXR0b25zIGhhdmUgZmluaXNoZWQgcmVuZGVyaW5nXG4gICAgICB0aGlzLiRuZXh0VGljayhmdW5jdGlvbiAoKSB7XG4gICAgICAgIHZhciBidG4gPSBfdGhpczIuZ2V0QnV0dG9ucygpLmZpbmQoZnVuY3Rpb24gKGVsKSB7XG4gICAgICAgICAgcmV0dXJuIHRvSW50ZWdlcihnZXRBdHRyKGVsLCAnYXJpYS1wb3NpbnNldCcpLCAwKSA9PT0gX3RoaXMyLmNvbXB1dGVkQ3VycmVudFBhZ2U7XG4gICAgICAgIH0pO1xuXG4gICAgICAgIGlmICghYXR0ZW1wdEZvY3VzKGJ0bikpIHtcbiAgICAgICAgICAvLyBGYWxsYmFjayBpZiBjdXJyZW50IHBhZ2UgaXMgbm90IGluIGJ1dHRvbiBsaXN0XG4gICAgICAgICAgX3RoaXMyLmZvY3VzRmlyc3QoKTtcbiAgICAgICAgfVxuICAgICAgfSk7XG4gICAgfSxcbiAgICBmb2N1c0ZpcnN0OiBmdW5jdGlvbiBmb2N1c0ZpcnN0KCkge1xuICAgICAgdmFyIF90aGlzMyA9IHRoaXM7XG5cbiAgICAgIC8vIFdlIGRvIHRoaXMgaW4gYCRuZXh0VGljaygpYCB0byBlbnN1cmUgYnV0dG9ucyBoYXZlIGZpbmlzaGVkIHJlbmRlcmluZ1xuICAgICAgdGhpcy4kbmV4dFRpY2soZnVuY3Rpb24gKCkge1xuICAgICAgICB2YXIgYnRuID0gX3RoaXMzLmdldEJ1dHRvbnMoKS5maW5kKGZ1bmN0aW9uIChlbCkge1xuICAgICAgICAgIHJldHVybiAhaXNEaXNhYmxlZChlbCk7XG4gICAgICAgIH0pO1xuXG4gICAgICAgIGF0dGVtcHRGb2N1cyhidG4pO1xuICAgICAgfSk7XG4gICAgfSxcbiAgICBmb2N1c0xhc3Q6IGZ1bmN0aW9uIGZvY3VzTGFzdCgpIHtcbiAgICAgIHZhciBfdGhpczQgPSB0aGlzO1xuXG4gICAgICAvLyBXZSBkbyB0aGlzIGluIGAkbmV4dFRpY2soKWAgdG8gZW5zdXJlIGJ1dHRvbnMgaGF2ZSBmaW5pc2hlZCByZW5kZXJpbmdcbiAgICAgIHRoaXMuJG5leHRUaWNrKGZ1bmN0aW9uICgpIHtcbiAgICAgICAgdmFyIGJ0biA9IF90aGlzNC5nZXRCdXR0b25zKCkucmV2ZXJzZSgpLmZpbmQoZnVuY3Rpb24gKGVsKSB7XG4gICAgICAgICAgcmV0dXJuICFpc0Rpc2FibGVkKGVsKTtcbiAgICAgICAgfSk7XG5cbiAgICAgICAgYXR0ZW1wdEZvY3VzKGJ0bik7XG4gICAgICB9KTtcbiAgICB9LFxuICAgIGZvY3VzUHJldjogZnVuY3Rpb24gZm9jdXNQcmV2KCkge1xuICAgICAgdmFyIF90aGlzNSA9IHRoaXM7XG5cbiAgICAgIC8vIFdlIGRvIHRoaXMgaW4gYCRuZXh0VGljaygpYCB0byBlbnN1cmUgYnV0dG9ucyBoYXZlIGZpbmlzaGVkIHJlbmRlcmluZ1xuICAgICAgdGhpcy4kbmV4dFRpY2soZnVuY3Rpb24gKCkge1xuICAgICAgICB2YXIgYnV0dG9ucyA9IF90aGlzNS5nZXRCdXR0b25zKCk7XG5cbiAgICAgICAgdmFyIGluZGV4ID0gYnV0dG9ucy5pbmRleE9mKGdldEFjdGl2ZUVsZW1lbnQoKSk7XG5cbiAgICAgICAgaWYgKGluZGV4ID4gMCAmJiAhaXNEaXNhYmxlZChidXR0b25zW2luZGV4IC0gMV0pKSB7XG4gICAgICAgICAgYXR0ZW1wdEZvY3VzKGJ1dHRvbnNbaW5kZXggLSAxXSk7XG4gICAgICAgIH1cbiAgICAgIH0pO1xuICAgIH0sXG4gICAgZm9jdXNOZXh0OiBmdW5jdGlvbiBmb2N1c05leHQoKSB7XG4gICAgICB2YXIgX3RoaXM2ID0gdGhpcztcblxuICAgICAgLy8gV2UgZG8gdGhpcyBpbiBgJG5leHRUaWNrKClgIHRvIGVuc3VyZSBidXR0b25zIGhhdmUgZmluaXNoZWQgcmVuZGVyaW5nXG4gICAgICB0aGlzLiRuZXh0VGljayhmdW5jdGlvbiAoKSB7XG4gICAgICAgIHZhciBidXR0b25zID0gX3RoaXM2LmdldEJ1dHRvbnMoKTtcblxuICAgICAgICB2YXIgaW5kZXggPSBidXR0b25zLmluZGV4T2YoZ2V0QWN0aXZlRWxlbWVudCgpKTtcblxuICAgICAgICBpZiAoaW5kZXggPCBidXR0b25zLmxlbmd0aCAtIDEgJiYgIWlzRGlzYWJsZWQoYnV0dG9uc1tpbmRleCArIDFdKSkge1xuICAgICAgICAgIGF0dGVtcHRGb2N1cyhidXR0b25zW2luZGV4ICsgMV0pO1xuICAgICAgICB9XG4gICAgICB9KTtcbiAgICB9XG4gIH0sXG4gIHJlbmRlcjogZnVuY3Rpb24gcmVuZGVyKGgpIHtcbiAgICB2YXIgX3RoaXM3ID0gdGhpcztcblxuICAgIHZhciBfc2FmZVZ1ZUluc3RhbmNlID0gc2FmZVZ1ZUluc3RhbmNlKHRoaXMpLFxuICAgICAgICBkaXNhYmxlZCA9IF9zYWZlVnVlSW5zdGFuY2UuZGlzYWJsZWQsXG4gICAgICAgIGxhYmVsUGFnZSA9IF9zYWZlVnVlSW5zdGFuY2UubGFiZWxQYWdlLFxuICAgICAgICBhcmlhTGFiZWwgPSBfc2FmZVZ1ZUluc3RhbmNlLmFyaWFMYWJlbCxcbiAgICAgICAgaXNOYXYgPSBfc2FmZVZ1ZUluc3RhbmNlLmlzTmF2LFxuICAgICAgICBudW1iZXJPZlBhZ2VzID0gX3NhZmVWdWVJbnN0YW5jZS5sb2NhbE51bWJlck9mUGFnZXMsXG4gICAgICAgIGN1cnJlbnRQYWdlID0gX3NhZmVWdWVJbnN0YW5jZS5jb21wdXRlZEN1cnJlbnRQYWdlO1xuXG4gICAgdmFyIHBhZ2VOdW1iZXJzID0gdGhpcy5wYWdlTGlzdC5tYXAoZnVuY3Rpb24gKHApIHtcbiAgICAgIHJldHVybiBwLm51bWJlcjtcbiAgICB9KTtcbiAgICB2YXIgX3RoaXMkcGFnaW5hdGlvblBhcmFtMiA9IHRoaXMucGFnaW5hdGlvblBhcmFtcyxcbiAgICAgICAgc2hvd0ZpcnN0RG90cyA9IF90aGlzJHBhZ2luYXRpb25QYXJhbTIuc2hvd0ZpcnN0RG90cyxcbiAgICAgICAgc2hvd0xhc3REb3RzID0gX3RoaXMkcGFnaW5hdGlvblBhcmFtMi5zaG93TGFzdERvdHM7XG4gICAgdmFyIGZpbGwgPSB0aGlzLmFsaWduID09PSAnZmlsbCc7XG4gICAgdmFyICRidXR0b25zID0gW107IC8vIEhlbHBlciBmdW5jdGlvbiBhbmQgZmxhZ1xuXG4gICAgdmFyIGlzQWN0aXZlUGFnZSA9IGZ1bmN0aW9uIGlzQWN0aXZlUGFnZShwYWdlTnVtYmVyKSB7XG4gICAgICByZXR1cm4gcGFnZU51bWJlciA9PT0gY3VycmVudFBhZ2U7XG4gICAgfTtcblxuICAgIHZhciBub0N1cnJlbnRQYWdlID0gdGhpcy5jdXJyZW50UGFnZSA8IDE7IC8vIEZhY3RvcnkgZnVuY3Rpb24gZm9yIHByZXYvbmV4dC9maXJzdC9sYXN0IGJ1dHRvbnNcblxuICAgIHZhciBtYWtlRW5kQnRuID0gZnVuY3Rpb24gbWFrZUVuZEJ0bihsaW5rVG8sIGFyaWFMYWJlbCwgYnRuU2xvdCwgYnRuVGV4dCwgYnRuQ2xhc3MsIHBhZ2VUZXN0LCBrZXkpIHtcbiAgICAgIHZhciBpc0Rpc2FibGVkID0gZGlzYWJsZWQgfHwgaXNBY3RpdmVQYWdlKHBhZ2VUZXN0KSB8fCBub0N1cnJlbnRQYWdlIHx8IGxpbmtUbyA8IDEgfHwgbGlua1RvID4gbnVtYmVyT2ZQYWdlcztcbiAgICAgIHZhciBwYWdlTnVtYmVyID0gbGlua1RvIDwgMSA/IDEgOiBsaW5rVG8gPiBudW1iZXJPZlBhZ2VzID8gbnVtYmVyT2ZQYWdlcyA6IGxpbmtUbztcbiAgICAgIHZhciBzY29wZSA9IHtcbiAgICAgICAgZGlzYWJsZWQ6IGlzRGlzYWJsZWQsXG4gICAgICAgIHBhZ2U6IHBhZ2VOdW1iZXIsXG4gICAgICAgIGluZGV4OiBwYWdlTnVtYmVyIC0gMVxuICAgICAgfTtcbiAgICAgIHZhciAkYnRuQ29udGVudCA9IF90aGlzNy5ub3JtYWxpemVTbG90KGJ0blNsb3QsIHNjb3BlKSB8fCB0b1N0cmluZyhidG5UZXh0KSB8fCBoKCk7XG4gICAgICB2YXIgJGlubmVyID0gaChpc0Rpc2FibGVkID8gJ3NwYW4nIDogaXNOYXYgPyBCTGluayA6ICdidXR0b24nLCB7XG4gICAgICAgIHN0YXRpY0NsYXNzOiAncGFnZS1saW5rJyxcbiAgICAgICAgY2xhc3M6IHtcbiAgICAgICAgICAnZmxleC1ncm93LTEnOiAhaXNOYXYgJiYgIWlzRGlzYWJsZWQgJiYgZmlsbFxuICAgICAgICB9LFxuICAgICAgICBwcm9wczogaXNEaXNhYmxlZCB8fCAhaXNOYXYgPyB7fSA6IF90aGlzNy5saW5rUHJvcHMobGlua1RvKSxcbiAgICAgICAgYXR0cnM6IHtcbiAgICAgICAgICByb2xlOiBpc05hdiA/IG51bGwgOiAnbWVudWl0ZW0nLFxuICAgICAgICAgIHR5cGU6IGlzTmF2IHx8IGlzRGlzYWJsZWQgPyBudWxsIDogJ2J1dHRvbicsXG4gICAgICAgICAgdGFiaW5kZXg6IGlzRGlzYWJsZWQgfHwgaXNOYXYgPyBudWxsIDogJy0xJyxcbiAgICAgICAgICAnYXJpYS1sYWJlbCc6IGFyaWFMYWJlbCxcbiAgICAgICAgICAnYXJpYS1jb250cm9scyc6IHNhZmVWdWVJbnN0YW5jZShfdGhpczcpLmFyaWFDb250cm9scyB8fCBudWxsLFxuICAgICAgICAgICdhcmlhLWRpc2FibGVkJzogaXNEaXNhYmxlZCA/ICd0cnVlJyA6IG51bGxcbiAgICAgICAgfSxcbiAgICAgICAgb246IGlzRGlzYWJsZWQgPyB7fSA6IHtcbiAgICAgICAgICAnIWNsaWNrJzogZnVuY3Rpb24gY2xpY2soZXZlbnQpIHtcbiAgICAgICAgICAgIF90aGlzNy5vbkNsaWNrKGV2ZW50LCBsaW5rVG8pO1xuICAgICAgICAgIH0sXG4gICAgICAgICAga2V5ZG93bjogb25TcGFjZUtleVxuICAgICAgICB9XG4gICAgICB9LCBbJGJ0bkNvbnRlbnRdKTtcbiAgICAgIHJldHVybiBoKCdsaScsIHtcbiAgICAgICAga2V5OiBrZXksXG4gICAgICAgIHN0YXRpY0NsYXNzOiAncGFnZS1pdGVtJyxcbiAgICAgICAgY2xhc3M6IFt7XG4gICAgICAgICAgZGlzYWJsZWQ6IGlzRGlzYWJsZWQsXG4gICAgICAgICAgJ2ZsZXgtZmlsbCc6IGZpbGwsXG4gICAgICAgICAgJ2QtZmxleCc6IGZpbGwgJiYgIWlzTmF2ICYmICFpc0Rpc2FibGVkXG4gICAgICAgIH0sIGJ0bkNsYXNzXSxcbiAgICAgICAgYXR0cnM6IHtcbiAgICAgICAgICByb2xlOiBpc05hdiA/IG51bGwgOiAncHJlc2VudGF0aW9uJyxcbiAgICAgICAgICAnYXJpYS1oaWRkZW4nOiBpc0Rpc2FibGVkID8gJ3RydWUnIDogbnVsbFxuICAgICAgICB9XG4gICAgICB9LCBbJGlubmVyXSk7XG4gICAgfTsgLy8gRWxsaXBzaXMgZmFjdG9yeVxuXG5cbiAgICB2YXIgbWFrZUVsbGlwc2lzID0gZnVuY3Rpb24gbWFrZUVsbGlwc2lzKGlzTGFzdCkge1xuICAgICAgcmV0dXJuIGgoJ2xpJywge1xuICAgICAgICBzdGF0aWNDbGFzczogJ3BhZ2UtaXRlbScsXG4gICAgICAgIGNsYXNzOiBbJ2Rpc2FibGVkJywgJ2J2LWQteHMtZG93bi1ub25lJywgZmlsbCA/ICdmbGV4LWZpbGwnIDogJycsIF90aGlzNy5lbGxpcHNpc0NsYXNzXSxcbiAgICAgICAgYXR0cnM6IHtcbiAgICAgICAgICByb2xlOiAnc2VwYXJhdG9yJ1xuICAgICAgICB9LFxuICAgICAgICBrZXk6IFwiZWxsaXBzaXMtXCIuY29uY2F0KGlzTGFzdCA/ICdsYXN0JyA6ICdmaXJzdCcpXG4gICAgICB9LCBbaCgnc3BhbicsIHtcbiAgICAgICAgc3RhdGljQ2xhc3M6ICdwYWdlLWxpbmsnXG4gICAgICB9LCBbX3RoaXM3Lm5vcm1hbGl6ZVNsb3QoU0xPVF9OQU1FX0VMTElQU0lTX1RFWFQpIHx8IHRvU3RyaW5nKF90aGlzNy5lbGxpcHNpc1RleHQpIHx8IGgoKV0pXSk7XG4gICAgfTsgLy8gUGFnZSBidXR0b24gZmFjdG9yeVxuXG5cbiAgICB2YXIgbWFrZVBhZ2VCdXR0b24gPSBmdW5jdGlvbiBtYWtlUGFnZUJ1dHRvbihwYWdlLCBpZHgpIHtcbiAgICAgIHZhciBwYWdlTnVtYmVyID0gcGFnZS5udW1iZXI7XG4gICAgICB2YXIgYWN0aXZlID0gaXNBY3RpdmVQYWdlKHBhZ2VOdW1iZXIpICYmICFub0N1cnJlbnRQYWdlOyAvLyBBY3RpdmUgcGFnZSB3aWxsIGhhdmUgdGFiaW5kZXggb2YgMCwgb3IgaWYgbm8gY3VycmVudCBwYWdlIGFuZCBmaXJzdCBwYWdlIGJ1dHRvblxuXG4gICAgICB2YXIgdGFiSW5kZXggPSBkaXNhYmxlZCA/IG51bGwgOiBhY3RpdmUgfHwgbm9DdXJyZW50UGFnZSAmJiBpZHggPT09IDAgPyAnMCcgOiAnLTEnO1xuICAgICAgdmFyIGF0dHJzID0ge1xuICAgICAgICByb2xlOiBpc05hdiA/IG51bGwgOiAnbWVudWl0ZW1yYWRpbycsXG4gICAgICAgIHR5cGU6IGlzTmF2IHx8IGRpc2FibGVkID8gbnVsbCA6ICdidXR0b24nLFxuICAgICAgICAnYXJpYS1kaXNhYmxlZCc6IGRpc2FibGVkID8gJ3RydWUnIDogbnVsbCxcbiAgICAgICAgJ2FyaWEtY29udHJvbHMnOiBzYWZlVnVlSW5zdGFuY2UoX3RoaXM3KS5hcmlhQ29udHJvbHMgfHwgbnVsbCxcbiAgICAgICAgJ2FyaWEtbGFiZWwnOiBoYXNQcm9wRnVuY3Rpb24obGFiZWxQYWdlKSA/XG4gICAgICAgIC8qIGlzdGFuYnVsIGlnbm9yZSBuZXh0ICovXG4gICAgICAgIGxhYmVsUGFnZShwYWdlTnVtYmVyKSA6IFwiXCIuY29uY2F0KGlzRnVuY3Rpb24obGFiZWxQYWdlKSA/IGxhYmVsUGFnZSgpIDogbGFiZWxQYWdlLCBcIiBcIikuY29uY2F0KHBhZ2VOdW1iZXIpLFxuICAgICAgICAnYXJpYS1jaGVja2VkJzogaXNOYXYgPyBudWxsIDogYWN0aXZlID8gJ3RydWUnIDogJ2ZhbHNlJyxcbiAgICAgICAgJ2FyaWEtY3VycmVudCc6IGlzTmF2ICYmIGFjdGl2ZSA/ICdwYWdlJyA6IG51bGwsXG4gICAgICAgICdhcmlhLXBvc2luc2V0JzogaXNOYXYgPyBudWxsIDogcGFnZU51bWJlcixcbiAgICAgICAgJ2FyaWEtc2V0c2l6ZSc6IGlzTmF2ID8gbnVsbCA6IG51bWJlck9mUGFnZXMsXG4gICAgICAgIC8vIEFSSUEgXCJyb3ZpbmcgdGFiaW5kZXhcIiBtZXRob2QgKGV4Y2VwdCBpbiBgaXNOYXZgIG1vZGUpXG4gICAgICAgIHRhYmluZGV4OiBpc05hdiA/IG51bGwgOiB0YWJJbmRleFxuICAgICAgfTtcbiAgICAgIHZhciBidG5Db250ZW50ID0gdG9TdHJpbmcoX3RoaXM3Lm1ha2VQYWdlKHBhZ2VOdW1iZXIpKTtcbiAgICAgIHZhciBzY29wZSA9IHtcbiAgICAgICAgcGFnZTogcGFnZU51bWJlcixcbiAgICAgICAgaW5kZXg6IHBhZ2VOdW1iZXIgLSAxLFxuICAgICAgICBjb250ZW50OiBidG5Db250ZW50LFxuICAgICAgICBhY3RpdmU6IGFjdGl2ZSxcbiAgICAgICAgZGlzYWJsZWQ6IGRpc2FibGVkXG4gICAgICB9O1xuICAgICAgdmFyICRpbm5lciA9IGgoZGlzYWJsZWQgPyAnc3BhbicgOiBpc05hdiA/IEJMaW5rIDogJ2J1dHRvbicsIHtcbiAgICAgICAgcHJvcHM6IGRpc2FibGVkIHx8ICFpc05hdiA/IHt9IDogX3RoaXM3LmxpbmtQcm9wcyhwYWdlTnVtYmVyKSxcbiAgICAgICAgc3RhdGljQ2xhc3M6ICdwYWdlLWxpbmsnLFxuICAgICAgICBjbGFzczoge1xuICAgICAgICAgICdmbGV4LWdyb3ctMSc6ICFpc05hdiAmJiAhZGlzYWJsZWQgJiYgZmlsbFxuICAgICAgICB9LFxuICAgICAgICBhdHRyczogYXR0cnMsXG4gICAgICAgIG9uOiBkaXNhYmxlZCA/IHt9IDoge1xuICAgICAgICAgICchY2xpY2snOiBmdW5jdGlvbiBjbGljayhldmVudCkge1xuICAgICAgICAgICAgX3RoaXM3Lm9uQ2xpY2soZXZlbnQsIHBhZ2VOdW1iZXIpO1xuICAgICAgICAgIH0sXG4gICAgICAgICAga2V5ZG93bjogb25TcGFjZUtleVxuICAgICAgICB9XG4gICAgICB9LCBbX3RoaXM3Lm5vcm1hbGl6ZVNsb3QoU0xPVF9OQU1FX1BBR0UsIHNjb3BlKSB8fCBidG5Db250ZW50XSk7XG4gICAgICByZXR1cm4gaCgnbGknLCB7XG4gICAgICAgIHN0YXRpY0NsYXNzOiAncGFnZS1pdGVtJyxcbiAgICAgICAgY2xhc3M6IFt7XG4gICAgICAgICAgZGlzYWJsZWQ6IGRpc2FibGVkLFxuICAgICAgICAgIGFjdGl2ZTogYWN0aXZlLFxuICAgICAgICAgICdmbGV4LWZpbGwnOiBmaWxsLFxuICAgICAgICAgICdkLWZsZXgnOiBmaWxsICYmICFpc05hdiAmJiAhZGlzYWJsZWRcbiAgICAgICAgfSwgcGFnZS5jbGFzc2VzLCBfdGhpczcucGFnZUNsYXNzXSxcbiAgICAgICAgYXR0cnM6IHtcbiAgICAgICAgICByb2xlOiBpc05hdiA/IG51bGwgOiAncHJlc2VudGF0aW9uJ1xuICAgICAgICB9LFxuICAgICAgICBrZXk6IFwicGFnZS1cIi5jb25jYXQocGFnZU51bWJlcilcbiAgICAgIH0sIFskaW5uZXJdKTtcbiAgICB9OyAvLyBHb3RvIGZpcnN0IHBhZ2UgYnV0dG9uXG4gICAgLy8gRG9uJ3QgcmVuZGVyIGJ1dHRvbiB3aGVuIGBoaWRlR290b0VuZEJ1dHRvbnNgIG9yIGBmaXJzdE51bWJlcmAgaXMgc2V0XG5cblxuICAgIHZhciAkZmlyc3RQYWdlQnRuID0gaCgpO1xuXG4gICAgaWYgKCF0aGlzLmZpcnN0TnVtYmVyICYmICF0aGlzLmhpZGVHb3RvRW5kQnV0dG9ucykge1xuICAgICAgJGZpcnN0UGFnZUJ0biA9IG1ha2VFbmRCdG4oMSwgdGhpcy5sYWJlbEZpcnN0UGFnZSwgU0xPVF9OQU1FX0ZJUlNUX1RFWFQsIHRoaXMuZmlyc3RUZXh0LCB0aGlzLmZpcnN0Q2xhc3MsIDEsICdwYWdpbmF0aW9uLWdvdG8tZmlyc3QnKTtcbiAgICB9XG5cbiAgICAkYnV0dG9ucy5wdXNoKCRmaXJzdFBhZ2VCdG4pOyAvLyBHb3RvIHByZXZpb3VzIHBhZ2UgYnV0dG9uXG5cbiAgICAkYnV0dG9ucy5wdXNoKG1ha2VFbmRCdG4oY3VycmVudFBhZ2UgLSAxLCB0aGlzLmxhYmVsUHJldlBhZ2UsIFNMT1RfTkFNRV9QUkVWX1RFWFQsIHRoaXMucHJldlRleHQsIHRoaXMucHJldkNsYXNzLCAxLCAncGFnaW5hdGlvbi1nb3RvLXByZXYnKSk7IC8vIFNob3cgZmlyc3QgKDEpIGJ1dHRvbj9cblxuICAgICRidXR0b25zLnB1c2godGhpcy5maXJzdE51bWJlciAmJiBwYWdlTnVtYmVyc1swXSAhPT0gMSA/IG1ha2VQYWdlQnV0dG9uKHtcbiAgICAgIG51bWJlcjogMVxuICAgIH0sIDApIDogaCgpKTsgLy8gRmlyc3QgZWxsaXBzaXNcblxuICAgICRidXR0b25zLnB1c2goc2hvd0ZpcnN0RG90cyA/IG1ha2VFbGxpcHNpcyhmYWxzZSkgOiBoKCkpOyAvLyBJbmRpdmlkdWFsIHBhZ2UgbGlua3NcblxuICAgIHRoaXMucGFnZUxpc3QuZm9yRWFjaChmdW5jdGlvbiAocGFnZSwgaWR4KSB7XG4gICAgICB2YXIgb2Zmc2V0ID0gc2hvd0ZpcnN0RG90cyAmJiBfdGhpczcuZmlyc3ROdW1iZXIgJiYgcGFnZU51bWJlcnNbMF0gIT09IDEgPyAxIDogMDtcbiAgICAgICRidXR0b25zLnB1c2gobWFrZVBhZ2VCdXR0b24ocGFnZSwgaWR4ICsgb2Zmc2V0KSk7XG4gICAgfSk7IC8vIExhc3QgZWxsaXBzaXNcblxuICAgICRidXR0b25zLnB1c2goc2hvd0xhc3REb3RzID8gbWFrZUVsbGlwc2lzKHRydWUpIDogaCgpKTsgLy8gU2hvdyBsYXN0IHBhZ2UgYnV0dG9uP1xuXG4gICAgJGJ1dHRvbnMucHVzaCh0aGlzLmxhc3ROdW1iZXIgJiYgcGFnZU51bWJlcnNbcGFnZU51bWJlcnMubGVuZ3RoIC0gMV0gIT09IG51bWJlck9mUGFnZXMgPyBtYWtlUGFnZUJ1dHRvbih7XG4gICAgICBudW1iZXI6IG51bWJlck9mUGFnZXNcbiAgICB9LCAtMSkgOiBoKCkpOyAvLyBHb3RvIG5leHQgcGFnZSBidXR0b25cblxuICAgICRidXR0b25zLnB1c2gobWFrZUVuZEJ0bihjdXJyZW50UGFnZSArIDEsIHRoaXMubGFiZWxOZXh0UGFnZSwgU0xPVF9OQU1FX05FWFRfVEVYVCwgdGhpcy5uZXh0VGV4dCwgdGhpcy5uZXh0Q2xhc3MsIG51bWJlck9mUGFnZXMsICdwYWdpbmF0aW9uLWdvdG8tbmV4dCcpKTsgLy8gR290byBsYXN0IHBhZ2UgYnV0dG9uXG4gICAgLy8gRG9uJ3QgcmVuZGVyIGJ1dHRvbiB3aGVuIGBoaWRlR290b0VuZEJ1dHRvbnNgIG9yIGBsYXN0TnVtYmVyYCBpcyBzZXRcblxuICAgIHZhciAkbGFzdFBhZ2VCdG4gPSBoKCk7XG5cbiAgICBpZiAoIXRoaXMubGFzdE51bWJlciAmJiAhdGhpcy5oaWRlR290b0VuZEJ1dHRvbnMpIHtcbiAgICAgICRsYXN0UGFnZUJ0biA9IG1ha2VFbmRCdG4obnVtYmVyT2ZQYWdlcywgdGhpcy5sYWJlbExhc3RQYWdlLCBTTE9UX05BTUVfTEFTVF9URVhULCB0aGlzLmxhc3RUZXh0LCB0aGlzLmxhc3RDbGFzcywgbnVtYmVyT2ZQYWdlcywgJ3BhZ2luYXRpb24tZ290by1sYXN0Jyk7XG4gICAgfVxuXG4gICAgJGJ1dHRvbnMucHVzaCgkbGFzdFBhZ2VCdG4pOyAvLyBBc3NlbWJsZSB0aGUgcGFnaW5hdGlvbiBidXR0b25zXG5cbiAgICB2YXIgJHBhZ2luYXRpb24gPSBoKCd1bCcsIHtcbiAgICAgIHN0YXRpY0NsYXNzOiAncGFnaW5hdGlvbicsXG4gICAgICBjbGFzczogWydiLXBhZ2luYXRpb24nLCB0aGlzLmJ0blNpemUsIHRoaXMuYWxpZ25tZW50LCB0aGlzLnN0eWxlQ2xhc3NdLFxuICAgICAgYXR0cnM6IHtcbiAgICAgICAgcm9sZTogaXNOYXYgPyBudWxsIDogJ21lbnViYXInLFxuICAgICAgICAnYXJpYS1kaXNhYmxlZCc6IGRpc2FibGVkID8gJ3RydWUnIDogJ2ZhbHNlJyxcbiAgICAgICAgJ2FyaWEtbGFiZWwnOiBpc05hdiA/IG51bGwgOiBhcmlhTGFiZWwgfHwgbnVsbFxuICAgICAgfSxcbiAgICAgIC8vIFdlIGRpc2FibGUga2V5Ym9hcmQgbGVmdC9yaWdodCBuYXYgd2hlbiBgPGItcGFnaW5hdGlvbi1uYXY+YFxuICAgICAgb246IGlzTmF2ID8ge30gOiB7XG4gICAgICAgIGtleWRvd246IHRoaXMuaGFuZGxlS2V5TmF2XG4gICAgICB9LFxuICAgICAgcmVmOiAndWwnXG4gICAgfSwgJGJ1dHRvbnMpOyAvLyBJZiB3ZSBhcmUgYDxiLXBhZ2luYXRpb24tbmF2PmAsIHdyYXAgaW4gYDxuYXY+YCB3cmFwcGVyXG5cbiAgICBpZiAoaXNOYXYpIHtcbiAgICAgIHJldHVybiBoKCduYXYnLCB7XG4gICAgICAgIGF0dHJzOiB7XG4gICAgICAgICAgJ2FyaWEtZGlzYWJsZWQnOiBkaXNhYmxlZCA/ICd0cnVlJyA6IG51bGwsXG4gICAgICAgICAgJ2FyaWEtaGlkZGVuJzogZGlzYWJsZWQgPyAndHJ1ZScgOiAnZmFsc2UnLFxuICAgICAgICAgICdhcmlhLWxhYmVsJzogaXNOYXYgPyBhcmlhTGFiZWwgfHwgbnVsbCA6IG51bGxcbiAgICAgICAgfVxuICAgICAgfSwgWyRwYWdpbmF0aW9uXSk7XG4gICAgfVxuXG4gICAgcmV0dXJuICRwYWdpbmF0aW9uO1xuICB9XG59KTtcblxuXG4vLyBXRUJQQUNLIEZPT1RFUiAvL1xuLy8gLi9ub2RlX21vZHVsZXMvYm9vdHN0cmFwLXZ1ZS9lc20vbWl4aW5zL3BhZ2luYXRpb24uanMiXSwibWFwcGluZ3MiOiJBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUVBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBRUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUVBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUVBO0FBRUE7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBRUE7QUFDQTtBQUNBIiwic291cmNlUm9vdCI6IiJ9\n//#
            sourceURL=webpack-internal:///./node_modules/bootstrap-vue/esm/mixins/pagination.js\n");
          </nav></b-pagination-nav
        ></b-pagination-nav
      ></b-pagination-nav
    ></b-pagination-nav
  ></b-pagination
>
