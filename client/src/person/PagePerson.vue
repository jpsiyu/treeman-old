<template>
    <div class="pagePerson">
        <Banner :activePage="swichPageAddState"></Banner>
        <div class="pagePerson-list">
            <Member v-for="info in members" :key="info.id" :info="info"/>
        </div>
        <PageAddPerson v-if="pageAddState" :swichPageAddState="swichPageAddState"></PageAddPerson>
    </div>
</template>

<script>
import { mapState } from "vuex";
import axios from "axios";
import Member from "./Member";
import Banner from "../common/Banner";
import PageAddPerson from "./PageAddPerson";
export default {
    name: "PagePerson",
    data: function() {
        return {
            pageAddState: false
        };
    },
    computed: {
        ...mapState({
            members: state => state.person.members
        })
    },
    methods: {
        swichPageAddState: function(b) {
            this.pageAddState = b;
        }
    },
    components: { Member, Banner, PageAddPerson },
    mounted: function() {
        axios.get("/api/allperson").then(response => {
            const data = response.data || [];
            this.$store.commit("initMembers", data);
        });
    }
};
</script>

<style scoped>
.pagePerson {
    position: relative;
}

.pagePerson-list {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    margin: auto;
    max-width: 768px;
}

@media screen and (max-width: 768px) {
    .pagePerson-list {
        justify-content: center;
    }
}
</style>
