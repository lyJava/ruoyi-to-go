<template>
	<div class="popup-result">
		<p class="title">最近5次运行时间</p>
		<ul class="popup-result-scroll">
			<template v-if="isShow">
				<li v-for="item in resultList" :key="item">{{ item }}</li>
			</template>
			<li v-else>计算结果中...</li>
		</ul>
	</div>
</template>

<script setup lang="ts">

interface DateArrays {
	[index: number]: number[];
}

const props = defineProps({
	ex: {
		type: String,
		required: true,
	},
});

const dayRule = ref<string>("");
const dayRuleSup = ref<any>([]);
const dateArr = reactive<DateArrays>([[], [], [], [], [], []]);
const resultList = ref<string[]>([]);
const isShow = ref(false);

const getOrderArr = (min: number, max: number): number[] => {
	return Array.from({ length: max - min + 1 }, (_, i) => min + i);
};

const getAssignArr = (rule: string): number[] => {
	return rule.split(",").map(Number).sort(compare);
};

const getAverageArr = (rule: string, limit: number): number[] => {
	const [minStr, stepStr] = rule.split("/");
	const min = Number(minStr);
	const step = Number(stepStr);
	const arr: number[] = [];
	for (let i = min; i <= limit; i += step) {
		arr.push(i);
	}
	return arr;
};

const getCycleArr = (
	rule: string,
	limit: number,
	status: boolean
): number[] => {
	const [startStr, endStr] = rule.split("-");
	let min = Number(startStr);
	let max = Number(endStr);
	const arr: number[] = [];

	if (min > max) {
		max += limit;
	}

	for (let i = min; i <= max; i++) {
		let value = i % limit;
		if (value === 0 && !status) {
			value = limit;
		}
		arr.push(value);
	}

	return arr.sort(compare);
};

// 日期处理方法
const formatDate = (value: Date | number, type?: "week"): string | number => {
	const time = typeof value === "number" ? new Date(value) : value;
	const Y = time.getFullYear();
	const M = time.getMonth() + 1;
	const D = time.getDate();
	const h = time.getHours();
	const m = time.getMinutes();
	const s = time.getSeconds();
	const week = time.getDay() + 1; // Vue3 中周日为0，这里保持原逻辑

	if (type === "week") return week;
	return `${Y}-${M.toString().padStart(2, "0")}-${D.toString().padStart(
		2,
		"0"
	)} ${h.toString().padStart(2, "0")}:${m.toString().padStart(2, "0")}:${s
		.toString()
		.padStart(2, "0")}`;
};

const checkDate = (value: string): boolean => {
	const time = new Date(value);
	return time.toString() !== "Invalid Date" && formatDate(time) === value;
};

// 时间计算核心方法
const getIndex = (arr: number[], value: number): number => {
	if (value <= arr[0] || value > arr[arr.length - 1]) return 0;
	for (let i = 0; i < arr.length - 1; i++) {
		if (value > arr[i] && value <= arr[i + 1]) return i + 1;
	}
	return 0;
};

const expressionChange = () => {
	isShow.value = false;
	const ruleArr = props.ex.split(" ");

	// 初始化时间数组
	const getTimeArrays = () => {
		getSecondArr(ruleArr[0]);
		getMinArr(ruleArr[1]);
		getHourArr(ruleArr[2]);
		getDayArr(ruleArr[3]);
		getMonthArr(ruleArr[4]);
		getYearArr(ruleArr[6], new Date().getFullYear());
		getWeekArr(ruleArr[5]);
	};

	getTimeArrays();
	// 用于记录进入循环的次数
	let nums = 0;
	// 用于暂时存符号时间规则结果的数组
	let resultArr = [];
	let nTime = new Date();
	let nYear = nTime.getFullYear();
	let nMonth = nTime.getMonth() + 1;
	let nDay = nTime.getDate();
	let nHour = nTime.getHours();
	let nMin = nTime.getMinutes();
	let nSecond = nTime.getSeconds();

	// 将获取到的数组赋值-方便使用
	let sDate = dateArr[0];
	let mDate = dateArr[1];
	let hDate = dateArr[2];
	let DDate = dateArr[3];
	let MDate = dateArr[4];
	let YDate = dateArr[5];
	// 获取当前时间在数组中的索引
	let sIdx = getIndex(sDate, nSecond);
	let mIdx = getIndex(mDate, nMin);
	let hIdx = getIndex(hDate, nHour);
	let DIdx = getIndex(DDate, nDay);
	let MIdx = getIndex(MDate, nMonth);
	let YIdx = getIndex(YDate, nYear);
	// 重置月日时分秒的函数(后面用的比较多)
	const resetSecond = function () {
		sIdx = 0;
		nSecond = sDate[sIdx];
	};
	const resetMin = function () {
		mIdx = 0;
		nMin = mDate[mIdx];
		resetSecond();
	};
	const resetHour = function () {
		hIdx = 0;
		nHour = hDate[hIdx];
		resetMin();
	};
	const resetDay = function () {
		DIdx = 0;
		nDay = DDate[DIdx];
		resetHour();
	};
	const resetMonth = function () {
		MIdx = 0;
		nMonth = MDate[MIdx];
		resetDay();
	};
	// 如果当前年份不为数组中当前值
	if (nYear !== YDate[YIdx]) {
		resetMonth();
	}
	// 如果当前月份不为数组中当前值
	if (nMonth !== MDate[MIdx]) {
		resetDay();
	}
	// 如果当前“日”不为数组中当前值
	if (nDay !== DDate[DIdx]) {
		resetHour();
	}
	// 如果当前“时”不为数组中当前值
	if (nHour !== hDate[hIdx]) {
		resetMin();
	}
	// 如果当前“分”不为数组中当前值
	if (nMin !== mDate[mIdx]) {
		resetSecond();
	}

	// 循环年份数组
	goYear: for (let Yi = YIdx; Yi < YDate.length; Yi++) {
		let YY = YDate[Yi];
		// 如果到达最大值时
		if (nMonth > MDate[MDate.length - 1]) {
			resetMonth();
			continue;
		}
		// 循环月份数组
		goMonth: for (let Mi = MIdx; Mi < MDate.length; Mi++) {
			// 赋值、方便后面运算
			let MM: any = MDate[Mi];
			MM = MM < 10 ? "0" + MM : MM;
			// 如果到达最大值时
			if (nDay > DDate[DDate.length - 1]) {
				resetDay();
				if (Mi == MDate.length - 1) {
					resetMonth();
					continue goYear;
				}
				continue;
			}
			// 循环日期数组
			goDay: for (let Di = DIdx; Di < DDate.length; Di++) {
				// 赋值、方便后面运算
				let DD: any = DDate[Di];
				let thisDD = DD < 10 ? "0" + DD : DD;

				// 如果到达最大值时
				if (nHour > hDate[hDate.length - 1]) {
					resetHour();
					if (Di == DDate.length - 1) {
						resetDay();
						if (Mi == MDate.length - 1) {
							resetMonth();
							continue goYear;
						}
						continue goMonth;
					}
					continue;
				}

				// 判断日期的合法性，不合法的话也是跳出当前循环
				if (
					checkDate(YY + "-" + MM + "-" + thisDD + " 00:00:00") !==
						true &&
					dayRule.value !== "workDay" &&
					dayRule.value !== "lastWeek" &&
					dayRule.value !== "lastDay"
				) {
					resetDay();
					continue goMonth;
				}
				// 如果日期规则中有值时
				if (dayRule.value == "lastDay") {
					// 如果不是合法日期则需要将前将日期调到合法日期即月末最后一天

					if (
						checkDate(
							YY + "-" + MM + "-" + thisDD + " 00:00:00"
						) !== true
					) {
						while (
							DD > 0 &&
							checkDate(
								YY + "-" + MM + "-" + thisDD + " 00:00:00"
							) !== true
						) {
							DD--;

							thisDD = DD < 10 ? "0" + DD : DD;
						}
					}
				} else if (dayRule.value == "workDay") {
					// 校验并调整如果是2月30号这种日期传进来时需调整至正常月底
					if (
						checkDate(
							YY + "-" + MM + "-" + thisDD + " 00:00:00"
						) !== true
					) {
						while (
							DD > 0 &&
							checkDate(
								YY + "-" + MM + "-" + thisDD + " 00:00:00"
							) !== true
						) {
							DD--;
							thisDD = DD < 10 ? "0" + DD : DD;
						}
					}
					// 获取达到条件的日期是星期X
					let thisWeek = formatDate(
						new Date(YY + "-" + MM + "-" + thisDD + " 00:00:00"),
						"week"
					);
					// 当星期日时
					if (thisWeek == 1) {
						// 先找下一个日，并判断是否为月底
						DD++;
						thisDD = DD < 10 ? "0" + DD : DD;
						// 判断下一日已经不是合法日期
						if (
							checkDate(
								YY + "-" + MM + "-" + thisDD + " 00:00:00"
							) !== true
						) {
							DD -= 3;
						}
					} else if (thisWeek == 7) {
						// 当星期6时只需判断不是1号就可进行操作
						if (dayRuleSup.value !== 1) {
							DD--;
						} else {
							DD += 2;
						}
					}
				} else if (dayRule.value == "weekDay") {
					// 如果指定了是星期几
					// 获取当前日期是属于星期几
					let thisWeek = formatDate(
						new Date(YY + "-" + MM + "-" + DD + " 00:00:00"),
						"week"
					);
					// 校验当前星期是否在星期池（dayRuleSup）中
					if (dayRuleSup.value.indexOf(thisWeek) < 0) {
						// 如果到达最大值时
						if (Di == DDate.length - 1) {
							resetDay();
							if (Mi == MDate.length - 1) {
								resetMonth();
								continue goYear;
							}
							continue goMonth;
						}
						continue;
					}
				} else if (dayRule.value == "assWeek") {
					// 如果指定了是第几周的星期几
					// 获取每月1号是属于星期几
					let thisWeek = formatDate(
						new Date(YY + "-" + MM + "-" + DD + " 00:00:00"),
						"week"
					);
					if (dayRuleSup.value[1].value >= thisWeek) {
						DD =
							(dayRuleSup.value[0] - 1) * 7 +
							dayRuleSup.value[1] -
							+thisWeek
							1;
					} else {
						DD =
							dayRuleSup.value[0] * 7 +
							dayRuleSup.value[1] -
							+thisWeek
							1;
					}
				} else if (dayRule.value == "lastWeek") {
					// 如果指定了每月最后一个星期几
					// 校验并调整如果是2月30号这种日期传进来时需调整至正常月底
					if (
						checkDate(
							YY + "-" + MM + "-" + thisDD + " 00:00:00"
						) !== true
					) {
						while (
							DD > 0 &&
							checkDate(
								YY + "-" + MM + "-" + thisDD + " 00:00:00"
							) !== true
						) {
							DD--;
							thisDD = DD < 10 ? "0" + DD : DD;
						}
					}
					// 获取月末最后一天是星期几
					let thisWeek = formatDate(
						new Date(YY + "-" + MM + "-" + thisDD + " 00:00:00"),
						"week"
					) as number;
					// 找到要求中最近的那个星期几
					if (dayRuleSup.value < thisWeek) {
						DD -= thisWeek - dayRuleSup.value;
					} else if (dayRuleSup.value > thisWeek) {
						DD -= 7 - (dayRuleSup.value - thisWeek);
					}
				}
				// 判断时间值是否小于10置换成“05”这种格式
				DD = DD < 10 ? "0" + DD : DD;

				// 循环“时”数组
				goHour: for (let hi = hIdx; hi < hDate.length; hi++) {
					let hh = hDate[hi] < 10 ? "0" + hDate[hi] : hDate[hi];

					// 如果到达最大值时
					if (nMin > mDate[mDate.length - 1]) {
						resetMin();
						if (hi == hDate.length - 1) {
							resetHour();
							if (Di == DDate.length - 1) {
								resetDay();
								if (Mi == MDate.length - 1) {
									resetMonth();
									continue goYear;
								}
								continue goMonth;
							}
							continue goDay;
						}
						continue;
					}
					// 循环"分"数组
					goMin: for (let mi = mIdx; mi < mDate.length; mi++) {
						let mm = mDate[mi] < 10 ? "0" + mDate[mi] : mDate[mi];

						// 如果到达最大值时
						if (nSecond > sDate[sDate.length - 1]) {
							resetSecond();
							if (mi == mDate.length - 1) {
								resetMin();
								if (hi == hDate.length - 1) {
									resetHour();
									if (Di == DDate.length - 1) {
										resetDay();
										if (Mi == MDate.length - 1) {
											resetMonth();
											continue goYear;
										}
										continue goMonth;
									}
									continue goDay;
								}
								continue goHour;
							}
							continue;
						}
						// 循环"秒"数组
						goSecond: for (
							let si = sIdx;
							si <= sDate.length - 1;
							si++
						) {
							let ss =
								sDate[si] < 10 ? "0" + sDate[si] : sDate[si];
							// 添加当前时间（时间合法性在日期循环时已经判断）
							if (MM !== "00" && DD !== "00") {
								resultArr.push(
									YY +
										"-" +
										MM +
										"-" +
										DD +
										" " +
										hh +
										":" +
										mm +
										":" +
										ss
								);
								nums++;
							}
							// 如果条数满了就退出循环
							if (nums == 5) break goYear;
							// 如果到达最大值时
							if (si == sDate.length - 1) {
								resetSecond();
								if (mi == mDate.length - 1) {
									resetMin();
									if (hi == hDate.length - 1) {
										resetHour();
										if (Di == DDate.length - 1) {
											resetDay();
											if (Mi == MDate.length - 1) {
												resetMonth();
												continue goYear;
											}
											continue goMonth;
										}
										continue goDay;
									}
									continue goHour;
								}
								continue goMin;
							}
						} 
					} 
				} 
			}
		} 
	}
	// 判断100年内的结果条数
	if (resultArr.length == 0) {
		resultList.value = ["没有达到条件的结果！"];
	} else {
		resultList.value = resultArr;
		if (resultArr.length !== 5) {
			resultList.value.push(
				"最近100年内只有上面" + resultArr.length + "条结果！"
			);
		}
	}
	// 计算完成-显示结果
	isShow.value = true;
};

// 获取"年"数组
const getYearArr = (rule: string, year: number) => {
	dateArr[5] = getOrderArr(year, year + 100);
	if (rule !== undefined) {
		if (rule.indexOf("-") >= 0) {
			dateArr[5] = getCycleArr(rule, year + 100, false);
		} else if (rule.indexOf("/") >= 0) {
			dateArr[5] = getAverageArr(rule, year + 100);
		} else if (rule !== "*") {
			dateArr[5] = getAssignArr(rule);
		}
	}
};
// 获取"月"数组
const getMonthArr = (rule: string) => {
	dateArr[4] = getOrderArr(1, 12);
	if (rule.indexOf("-") >= 0) {
		dateArr[4] = getCycleArr(rule, 12, false);
	} else if (rule.indexOf("/") >= 0) {
		dateArr[4] = getAverageArr(rule, 12);
	} else if (rule !== "*") {
		dateArr[4] = getAssignArr(rule);
	}
};
// 获取"日"数组-主要为日期规则
const getWeekArr = (rule: string) => {
	// 只有当日期规则的两个值均为“”时则表达日期是有选项的
	if (dayRule.value == "" && dayRuleSup.value == "") {
		if (rule.indexOf("-") >= 0) {
			dayRule.value = "weekDay";
			dayRuleSup.value = getCycleArr(rule, 7, false);
		} else if (rule.indexOf("#") >= 0) {
			dayRule.value = "assWeek";
			let matchRule = rule.match(/[0-9]{1}/g) as any;
			dayRuleSup.value = [Number(matchRule[1]!), Number(matchRule[0])];
			dateArr[3] = [1];
			if (dayRuleSup.value[1] == 7) {
				dayRuleSup.value[1] = 0;
			}
		} else if (rule.indexOf("L") >= 0) {
			dayRule.value = "lastWeek";
			dayRuleSup.value = Number(rule.match(/[0-9]{1,2}/g)?.[0]);
			dateArr[3] = [31];
			if (dayRuleSup.value == 7) {
				dayRuleSup.value = 0;
			}
		} else if (rule !== "*" && rule !== "?") {
			dayRule.value = "weekDay";
			dayRuleSup.value = getAssignArr(rule);
		}
	}
};
// 获取"日"数组-少量为日期规则
const getDayArr = (rule: string) => {
	dateArr[3] = getOrderArr(1, 31);
	dayRule.value = "";
	dayRuleSup.value = "";
	if (rule.indexOf("-") >= 0) {
		dateArr[3] = getCycleArr(rule, 31, false);
		dayRuleSup.value = "null";
	} else if (rule.indexOf("/") >= 0) {
		dateArr[3] = getAverageArr(rule, 31);
		dayRuleSup.value = "null";
	} else if (rule.indexOf("W") >= 0) {
		dayRule.value = "workDay";
		dayRuleSup.value = Number(rule?.match(/[0-9]{1,2}/g)?.[0]);
		dateArr[3] = [dayRuleSup.value];
	} else if (rule.indexOf("L") >= 0) {
		dayRule.value = "lastDay";
		dayRuleSup.value = "null";
		dateArr[3] = [31];
	} else if (rule !== "*" && rule !== "?") {
		dateArr[3] = getAssignArr(rule);
		dayRuleSup.value = "null";
	} else if (rule == "*") {
		dayRuleSup.value = "null";
	}
};

// 各个时间字段的处理方法
const getSecondArr = (rule: string) => {
	dateArr[0] = handleFieldRule(rule, 0, 59);
};

const getMinArr = (rule: string) => {
	dateArr[1] = handleFieldRule(rule, 0, 59);
};

const getHourArr = (rule: string) => {
	dateArr[2] = handleFieldRule(rule, 0, 23);
};

const handleFieldRule = (rule: string, min: number, max: number): number[] => {
	if (rule.includes("-")) {
		return getCycleArr(rule, max + 1, true);
	}
	if (rule.includes("/")) {
		return getAverageArr(rule, max);
	}
	if (rule !== "*") {
		return getAssignArr(rule);
	}
	return getOrderArr(min, max);
};

// 比较数字大小（用于Array.sort）
const compare = (value1: number, value2: number) => {
	if (value2 - value1 > 0) {
		return -1;
	} else {
		return 1;
	}
};

watch(() => props.ex, expressionChange);
onMounted(expressionChange);
</script>
