<template>
    <div class="record">
        <Banner :activePage="switchPageAddRecordState"></Banner>
    </div>
</template>

<script>
import Banner from "../common/Banner";
export default {
    data: function() {
        return {
            pageAddRecordState: false
        };
    },
    components: { Banner },
    methods: {
        switchPageAddRecordState: function(b) {
            this.pageAddRecordState = b;
        }
    }
};
</script>

<style scoped>
</style>

