// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace Aliyun.SDK.CCP.CCPClient.Models
{
    /**
     * list file request
     */
    public class OSSListFileRequest : TeaModel {
        [NameInMap("drive_id")]
        [Validation(Required=false, Pattern="[0-9]+")]
        public string DriveId { get; set; }

        [NameInMap("image_thumbnail_process")]
        [Validation(Required=false)]
        public string ImageThumbnailProcess { get; set; }

        [NameInMap("image_url_process")]
        [Validation(Required=false)]
        public string ImageUrlProcess { get; set; }

        [NameInMap("limit")]
        [Validation(Required=false, Pattern="[0-9]{1,3}")]
        public long Limit { get; set; }

        [NameInMap("marker")]
        [Validation(Required=false)]
        public string Marker { get; set; }

        [NameInMap("parent_file_path")]
        [Validation(Required=true)]
        public string ParentFilePath { get; set; }

        [NameInMap("share_id")]
        [Validation(Required=false, Pattern="[0-9]+")]
        public string ShareId { get; set; }

        [NameInMap("url_expire_sec")]
        [Validation(Required=false)]
        public long UrlExpireSec { get; set; }

        [NameInMap("video_thumbnail_process")]
        [Validation(Required=false)]
        public string VideoThumbnailProcess { get; set; }

    }

}
