<template>
	<el-image
		:src="`${realSrc}`"
		fit="cover"
		:style="`width:${realWidth};height:${realHeight};`"
		:preview-src-list="realSrcList"
	>
		<div slot="error" class="image-slot">
			<i class="picture-outline"></i>
		</div>
	</el-image>
</template>

<script lang="ts" setup>
import { isExternal } from "@/utils/validate";
const props = defineProps({
	src: {
		type: String,
		required: true,
	},
	width: {
		type: [Number, String],
		default: "",
	},
	height: {
		type: [Number, String],
		default: "",
	},
});

const realSrc = () => {
	let real_src = props.src.split(",")[0];
	if (isExternal(real_src)) {
		return real_src;
	}
	return import.meta.env.VUE_APP_BASE_API + real_src;
};
const realSrcList = () => {
	let real_src_list = props.src.split(",");
	let srcList: string[] = [];
	real_src_list.forEach((item) => {
		if (isExternal(item)) {
			return srcList.push(item);
		}
		return srcList.push(import.meta.env.VUE_APP_BASE_API + item);
	});
	return srcList;
};
const realWidth = () => {
	return typeof props.width == "string" ? props.width : `${props.width}px`;
};
const realHeight = () => {
	return typeof props.height == "string" ? props.height : `${props.height}px`;
};
console.log("触发了图标预览组件");
onMounted(() => {
    realSrc();
    realSrcList();
    realWidth();
    realHeight();
});

</script>

<style lang="scss" scoped>
.el-image {
	border-radius: 5px;
	background-color: #ebeef5;
	box-shadow: 0 0 5px 1px #ccc;
	:v-deep .el-image__inner {
		transition: all 0.3s;
		cursor: pointer;
		&:hover {
			transform: scale(1.2);
		}
	}
	:v-deep .image-slot {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		height: 100%;
		color: #909399;
		font-size: 30px;
	}
}
</style>
