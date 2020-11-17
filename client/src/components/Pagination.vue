<template>
<div>
    <nav aria-label="Pagination">
        <ul class="pagination mt-2 justify-content-center">
            <li :key="'prev'" :class="{'disabled':currentPage === 1 || totalPages === 0}" class="page-item previous-item">
                <button class="page-link" @click="currentPage--">Prev</button>
            </li>
            <li v-for="page in pageArray" :key="page" class="page-item" :class="{'active':(currentPage === page)}">
                <button class="page-link" @click="currentPage = page">{{page}}</button>
            </li>
            <li :key="'next'" :class="{'disabled':currentPage === totalPages || totalPages === 0}" class="page-item next-item">
                <button class="page-link" @click="currentPage++">Next</button>
            </li>
        </ul>
    </nav>
</div>
</template>

<script>
export default {
    name: "pagination",
    props: {
        method: {
            type: Function
        },
        numOfDocuments: Number,
        onPageChange: Function,
    },
    methods: {
        validatePageCount(count, total) {
            if (count > total) {
                count = total;
            }
            if (count % 2 == 0) { //ensure odd number
                count--;
            }
            return count;
        },

        generatePaginationPageArray(numOfDocuments, recordsPerPage) { // Fetch our pages in an array so we can iterate over it later to create the pagination list items.
            if (numOfDocuments == 0) { // No paging if no records.
                this.totalPages = 0;
                return;
            }

            const pageArray = [];
            this.totalPages = Math.ceil(numOfDocuments / recordsPerPage); // sets total page count for data
            let pageDisplayCount = this.validatePageCount(5, this.totalPages);
            const mid = Math.ceil(pageDisplayCount / 2); // sets mid of displayed pages
            let startIndex = (this.currentPage + 1) - mid; // set first page to be displayed

            if (this.totalPages <= recordsPerPage) {
                startIndex = 1;
                pageDisplayCount = this.totalPages;
            } else if (this.currentPage < mid) { // account for front half
                startIndex = 1;
            } else if ((this.currentPage + (mid - 1)) > this.totalPages) { //account for back half
                startIndex = this.totalPages - pageDisplayCount + 1;
            }
            for (let i = startIndex; i < startIndex + pageDisplayCount; i++) {
                pageArray.push(i);
            }
            return pageArray;
        },

        reset() {
            this.currentPage = 1;
            console.log("RESET PAGE" + this.currentPage);
            this.$emit('reset');
        }
    },
    data() {
        return {
            pageArray: [],
            currentPage: typeof this.$route.query.page == "undefined" ? 1 : parseInt(this.$route.query.page),
            totalPages: 0
        };
    },
    mounted() {},

    watch: { // Watch for data change in which page the user is currently on, call API to get new data when it changes.
        "numOfDocuments": function (val) {
            console.log(val);
            this.pageArray = this.generatePaginationPageArray(val, 5);
        },
        "currentPage": function (page) {
            console.log("Page Changed");
            this.onPageChange(page);
        }
    }
};
</script>
