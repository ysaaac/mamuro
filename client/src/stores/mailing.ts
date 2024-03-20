import {defineStore} from 'pinia';
import {MailSectionType} from "@/enums/MailSection";
import type {MailListType} from "@/types/MailList";
import {formatDate} from "@/utils/FormatDateTime";

export const useMailingStore = defineStore({
    id: 'mail',
    state: () => ({
        mailSectionType: MailSectionType.INBOX as MailSectionType,
        mailList: [] as MailListType[],
        totalMails: 0 as number
    }),
    actions: {
        setMailListData(data: MailListResponseType) {
            this.totalMails = data.hits.total.value
            this.mailList = data.hits.hits.map(mailResponse => {
                return {
                    mailId: mailResponse._id,
                    from: mailResponse._source.from,
                    subject: mailResponse._source.subject,
                    date: formatDate(mailResponse._source.date),
                    viewed: false,
                }
            })
            if (this.mailList.length >= 1) this.mailList[0].viewed = true;
        },
    },
});

interface MailResponseType {
    _id: string
    _score: number
    "@timestamp": string
    _source: {
        date: string
        from: string
        message_id: string
        subject: string
    }
}

interface HitsType {
    total: {
        value: number
    },
    hits: MailResponseType[]
}

interface MailListResponseType {
    hits: HitsType
}