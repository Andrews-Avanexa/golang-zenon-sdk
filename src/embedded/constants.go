package embedded

import (
	"regexp"
)

var fuseMinQsrAmount int = 10 * utils.oneQsr
var minPlasmaAmount int = 21000

// Pillar
var pillarRegisterZnnAmount int = 15000 * utils.oneZnn
var pillarRegisterQsrAmount int = 150000 * utils.oneQsr
var pillarNameMaxLength int = 40
var pillarNameRegExp, _ = regexp.Compile("^([a-zA-Z0-9]+[-._]?)*[a-zA-Z0-9]$")

// Sentinel
var sentinelRegisterZnnAmount int = 5000 * utils.oneZnn
var sentinelRegisterQsrAmount int = 50000 * utils.oneQsr

// Staking
var stakeTimeUnitSec int = 24 * 60 * 60
var stakeTimeMaxSec int = 12 * stakeTimeUnitSec
var stakeMinZnnAmount int = 1 * utils.oneZnn
var stakeUnitDurationName string = "days"

// Token
var tokenZtsIssueFeeInZnn int = 100 * utils.oneZnn
var tokenNameMaxLength int = 40
var tokenNameRegExp, _ = regexp.Compile("^([a-zA-Z0-9]+[-._]?)*[a-zA-Z0-9]$")
var tokenSymbolRegExp, _ = regexp.Compile("^[A-Z0-9]+$")
var tokenSymbolMaxLength int = 10

// var tokenSymbolExceptions []string = ["ZNN", "QSR"];
var tokenDomainRegExp, _ = regexp.Compile("^([A-Za-z0-9][A-Za-z0-9-]{0,61}[A-Za-z0-9]\\.)+[A-Za-z]{2,}$")

// Accelerator
var proposalUrlRegExp, _ = regexp.Compile("^[a-zA-Z0-9]{2,60}\\.[a-zA-Z]{1,6}([a-zA-Z0-9()@:%_\\+.~#?&/=-]{0,100})$")
var proposalDescriptionMaxLength int = 240
var proposalNameMaxLength int = 30

const proposalCreationCostInZnn int = 10
const proposalMaximumFundsInZnn int = 5000
const proposalMinimumFundsInZnn int = 10
const proposalVotingStatus int = 0
const proposalActiveStatus int = 1
const proposalPaidStatus int = 2
const proposalClosedStatus int = 3
