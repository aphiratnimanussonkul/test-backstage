import React from 'react';

import {
    Header,
    Page,
    pageTheme,
} from '@backstage/core';

const SearchCoursePage = () => {
    return <>
        <Page theme={pageTheme.home}>
            <Header
                title='ระบบค้นหาข้อมูลหลักสูตร'
                subtitle="เพื่อค้นหาข้อมูลหลักสูตรต่างๆภายในมหาลัย"
            ></Header>
        </Page>

    </>
}

export default SearchCoursePage;