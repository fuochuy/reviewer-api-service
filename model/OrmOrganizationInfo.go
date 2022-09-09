package model

//
//import "google.golang.org/genproto/googleapis/type/date"
//
//type OrmOrganizationInfo struct {
//	organizationId int64	`json:"organization_id"`
//
//	name	string `json:"name"`
//
//	displayName string	`json:"display_name	"`
//
//	description string 	`json:"description"`
//
//	address string `json:"address"`
//
//	createdAt date.Date `json:"created_at"`
//
//	expiredAt date.Date	`json:"expired_at"`
//
//	updatedAt date.Date	`json:"updated_at"`
//
//	avatar string `json:"avatar"`
//
//	searchType int64 `json:"search_type"`
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.TOPIC)
//	private TopicInfo topic;
//
//	private Map<String, Object> other = new HashMap<>();
//
//	@JsonAnyGetter
//	public Map<String, Object> any() {
//	return other;
//}
//
//	@JsonAnySetter
//	public void set(String name, Object value) {
//	other.put(name, value);
//}
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.RULES_COUNT)
//	private int rulesCount;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.USERS_COUNT)
//	private int usersCount;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.TOPICS_COUNT)
//	private int topicsCount;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.CURRENT_PACKAGE)
//	private String currentPackage;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.PACKAGE_NAME)
//	private String packageName;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.TOPICS)
//	private List<TopicInfo> topics;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.USERS)
//	private List<OrmUserInfo> users;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.REQUEST_HISTORY)
//	private List<UserRequests> requestHistory;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.DEPARTMENT)
//	private String department;
//
//	@JsonSerialize(using = CustomDateSerializer.class)
//	@JsonDeserialize(using = CustomDateDeSerializer.class)
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.PAID_START_TIME)
//	private Date paidStartTime;
//
//	@JsonSerialize(using = CustomDateSerializer.class)
//	@JsonDeserialize(using = CustomDateDeSerializer.class)
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.PAID_END_TIME)
//	private Date paidEndTime;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.PURPOSE)
//	private String purpose;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.TOTAL_DATA_OF_MONTH)
//	private TotalDataOfMonth totalDataOfMonth;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.LAST_ACCESSED)
//	private String lastAccessed;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.NUMBER_OF_ACCESSES)
//	private Long numberOfAccesses;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.MAX_ACCOUNT)
//	private Integer maxAccount;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.INDUSTRIES)
//	private List<Industry> industries = new ArrayList<>();
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.RULES)
//	private List<RuleInfo> ruleInfos = new ArrayList<>();
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.MARKETING_CHANNEL)
//	private String marketingChannel;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.ADMIN_EMAIL)
//	private String adminEmail;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.ADMIN_PHONE)
//	private String adminPhone;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.ADMIN_USERNAME)
//	private String adminUserName;
//
//	private String location;
//
//	private Integer type;
//
//	// 0: demo - 1: paid
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.IS_PAID)
//	private Integer isPaid;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.ADMIN_NAME)
//	private String adminName;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.IS_ADDITIONAL_SOURCE)
//	private Integer additionalSource;
//
//	@JsonProperty(value = Constant.OrmOrganizationInfo.JsonField.REPORT_INFO)
//	private ReportInfo reportInfo;
//}
