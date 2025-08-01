<template>
	<div class="upload-file">
		<el-upload
			multiple
			:action="uploadFileUrl"
			:before-upload="handleBeforeUpload"
			:file-list="fileList"
			:limit="limit"
			:on-error="handleUploadError"
			:on-exceed="handleExceed"
			:on-success="handleUploadSuccess"
			:show-file-list="false"
			:headers="headers"
			class="upload-file-uploader"
			ref="upload"
		>
			<!-- 上传按钮 -->
			<el-button type="primary">选取文件</el-button>
		</el-upload>
		<!-- 上传提示 -->
		<div class="el-upload__tip" v-if="showTip">
			请上传
			<template v-if="fileSize">
				大小不超过 <b style="color: #f56c6c">{{ fileSize }}MB</b>
			</template>
			<template v-if="fileType">
				格式为 <b style="color: #f56c6c">{{ fileType.join("/") }}</b>
			</template>
			的文件
		</div>
		<!-- 文件列表 -->
		<transition-group
			class="upload-file-list el-upload-list el-upload-list--text"
			name="el-fade-in-linear"
			tag="ul"
		>
			<li
				:key="file.uid"
				class="el-upload-list__item ele-upload-list__item-content"
				v-for="(file, index) in fileList"
			>
				<el-link
					:href="`${baseUrl}${file.url}`"
					:underline="false"
					target="_blank"
				>
					<span class="document">
						{{ getFileName(file.name) }}
					</span>
				</el-link>
				<div class="ele-upload-list__item-content-action">
					<el-link
						:underline="false"
						@click="handleDelete(index)"
						type="danger"
						>删除</el-link
					>
				</div>
			</li>
		</transition-group>
	</div>
</template>

<script setup lang="ts">
import { getToken } from "@/utils/auth";
import { useSafeInstance } from "@/utils/ruoyi";
import { UploadFile, UploadRawFile } from "element-plus";

const props = defineProps({
	modelValue: [String, Object, Array],
	limit: {
		type: Number,
		default: 5,
	},
	fileSize: {
		type: Number,
		default: 5,
	},
	fileType: {
		type: Array,
		default: ["doc", "xls", "ppt", "txt", "pdf"],
	},
	// 是否显示提示
	isShowTip: {
		type: Boolean,
		default: true,
	},
});

const proxy = useSafeInstance();
const number = ref<number>(0);
const uploadList = ref<UploadFile[]>([]);
const baseUrl = import.meta.env.VITE_APP_BASE_API;
const uploadFileUrl = import.meta.env.VITE_APP_BASE_API + "/common/upload";
const headers = { Authorization: "Bearer " + getToken() };
const fileList = ref<UploadFile[] | undefined>([]);
const showTip = computed(
	() => props.isShowTip && (props.fileType || props.fileSize)
);

watch(
	() => props.modelValue,
	(val) => {
		if (val) {
			let temp = 1;
			// 首先将值转为数组
			const list = Array.isArray(val)
				? val
				: (props.modelValue?.toString().split(","));
			// 然后将数组转为对象数组
			fileList.value = list?.map((item) => {
				if (typeof item === "string") {
					item = { name: item, url: item };
				}
				item.uid = item.uid || new Date().getTime() + temp++;
				return item;
			});
		} else {
			fileList.value = [];
			return [];
		}
	},
	{ deep: true, immediate: true }
);

// 上传前校检格式和大小
const handleBeforeUpload = (file: UploadRawFile) => {
	// 校检文件类型
	if (props.fileType.length) {
		let fileExtension = "";
		if (file.name.lastIndexOf(".") > -1) {
			fileExtension = file.name.slice(file.name.lastIndexOf(".") + 1);
		}
		const isTypeOk = props.fileType.some((type: any) => {
			if (file.type.indexOf(type) > -1) return true;
			if (fileExtension && fileExtension.indexOf(type) > -1) return true;
			return false;
		});
		if (!isTypeOk) {
			proxy.$modal.msgError(
				`文件格式不正确, 请上传${props.fileType.join("/")}格式文件!`
			);
			return false;
		}
	}
	// 校检文件大小
	if (props.fileSize) {
		const isLt = file.size / 1024 / 1024 < props.fileSize;
		if (!isLt) {
			proxy.$modal.msgError(`上传文件大小不能超过 ${props.fileSize} MB!`);
			return false;
		}
	}
	proxy.$modal.loading("正在上传文件，请稍候...");
	number.value++;
	return true;
};

// 文件个数超出
const handleExceed = () => {
	proxy.$modal.msgError(`上传文件数量不能超过 ${props.limit} 个!`);
};

// 上传失败
const handleUploadError = (err: Error) => {
	console.log("上传文件失败", err);
	proxy.$modal.msgError("上传文件失败");
};


const emit = defineEmits(["update:modelValue"]);

// 上传成功回调
const handleUploadSuccess = (uploadFile: UploadFile) => {
	uploadList.value.push(uploadFile);
	if (uploadList.value.length === number.value) {
		fileList.value = fileList.value!
			.filter((f) => f.url !== undefined)
			.concat(uploadList.value);
		uploadList.value = [];
		number.value = 0;
		emit("update:modelValue", listToString(fileList.value));
		proxy.$modal.closeLoading();
	}
};

// 删除文件
const handleDelete = (index: number) => {
	fileList.value?.splice(index, 1);
	emit("update:modelValue", listToString(fileList.value!));
};

// 获取文件名称
const getFileName = (name: string) => {
	if (name.lastIndexOf("/") > -1) {
		return name.slice(name.lastIndexOf("/") + 1);
	} else {
		return "";
	}
};

// 对象转成指定字符串分隔
const listToString = (list: any[], separator?: string) => {
	let strs = "";
	separator = separator || ",";
	for (let i in list) {
		if (undefined !== list[i].url) {
			strs += list[i].url + separator;
		}
	}
	return strs !== "" ? strs.substring(0, strs.length - 1) : "";
};
</script>

<style scoped lang="scss">
.upload-file-uploader {
	margin-bottom: 5px;
}
.upload-file-list .el-upload-list__item {
	border: 1px solid #e4e7ed;
	line-height: 2;
	margin-bottom: 10px;
	position: relative;
}
.upload-file-list .ele-upload-list__item-content {
	display: flex;
	justify-content: space-between;
	align-items: center;
	color: inherit;
}
.ele-upload-list__item-content-action .el-link {
	margin-right: 10px;
}
</style>
