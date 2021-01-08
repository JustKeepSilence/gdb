<template>
    <div style="height:100%">
        <iframe :src="url" style="width:100%;height: 100%;border:none">
        </iframe>
    </div>
</template>
<script>
import { getCookie } from '@/utils/cookie'
export default {
    name: 'Document',
    data(){
        return {
            url: `http://${getCookie("ip").split(":")[0]}:8087/`
        }
    },
    mounted(){
        document.querySelector(".el-main").style.backgroundColor = " #f0f2f5";
    }
}
</script>