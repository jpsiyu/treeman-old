<template>
    <div class="page-person">
        <Banner></Banner>
        <div class="members">
            <Member v-for="info in members" :key="info.id" :info="info"/>
        </div>
        <PageAdd v-if="pageAddOn"></PageAdd>
    </div>
</template>

<script>
import { mapState } from "vuex";
import Member from "./Member";
import Banner from "./Banner";
import PageAdd from "./PageAdd";
export default {
    name: "PagePerson",
    components: { Member, Banner, PageAdd },
    computed: {
        ...mapState({
            pageAddOn: state => state.person.pageAddOn,
            members: state => state.person.members
        })
    }
};
</script>

<style scoped>
.page-person{
    position: relative;
}

.members {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    margin: auto;
    max-width: 768px;
}

@media screen and (max-width: 768px) {
    .members {
        justify-content: center;
    }
}
</style>
